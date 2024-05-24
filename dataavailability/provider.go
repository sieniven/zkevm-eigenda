package dataavailability

import (
	"context"
	"encoding/binary"
	"fmt"
	"sync"

	disperser_rpc "github.com/Layr-Labs/eigenda/api/grpc/disperser"
	"github.com/Layr-Labs/eigenda/encoding/utils/codec"
	"github.com/ethereum/go-ethereum/common"
)

// DataAvailabilityProvider is the EigenDA backend manager that holds the DA implementation.
// It contains the implementation of SequenceSender and SequenceRetriever of the zkevm's
// dataavailability package.
type DataAvailabilityProvider struct {
	cfg    Config
	state  DAStorage
	client *DisperserClient
}

// Factory method for a new DataAvailibilityProvider instance
func NewDataProvider(cfg Config) *DataAvailabilityProvider {
	// Initialize in-memory DA storage
	s := DAStorage{
		inner:    map[common.Hash]BlobInfo{},
		da_inner: map[common.Hash]int{},
		mutex:    &sync.RWMutex{},
	}
	signer := MockBlobRequestSigner{}

	p := &DataAvailabilityProvider{
		cfg:    cfg,
		state:  s,
		client: NewDisperserClient(&cfg, signer),
	}
	return p
}

func (d *DataAvailabilityProvider) Init() error {
	return nil
}

func (d *DataAvailabilityProvider) PostSequence(ctx context.Context, batchesData [][]byte) ([]byte, error) {
	blobData := EncodeSequence(batchesData)

	// Blob serialization
	blobData = codec.ConvertByPaddingEmptyByte(blobData)

	// Send blob to EigenDA disperser
	_, idBytes, err := d.client.DisperseBlob(ctx, blobData, []uint8{})
	if err != nil {
		fmt.Println("failed to send blob to EigenDA disperser")
		return []byte{}, nil
	}

	fmt.Println("sent blob to EigenDA disperser")
	for {
		blobStatusReply, err := d.client.GetBlobStatus(ctx, idBytes)
		if err != nil {
			fmt.Printf("error getting blob status: %v\n", err)
			return nil, err
		}

		// Get blob status
		currStatus := blobStatusReply.GetStatus()
		if currStatus == disperser_rpc.BlobStatus_CONFIRMED {
			break
		}
	}
	return idBytes, nil
}

func (d *DataAvailabilityProvider) GetSequence(ctx context.Context, batchHashes []common.Hash, blobInfo BlobInfo) ([][]byte, error) {
	var batchesData [][]byte
	for _, hash := range batchHashes {
		batchData, err := d.GetBatchL2Data(ctx, hash)
		if err != nil {
			return nil, err
		}
		batchesData = append(batchesData, batchData)
	}
	reply, err := d.client.RetrieveBlob(ctx, blobInfo.BatchHeaderHash, blobInfo.BlobIndex)
	if err != nil {
		fmt.Printf("failed to retrieve blob: %v\n", err)
		return nil, err
	}
	batchesData = append(batchesData, reply.GetData())
	return batchesData, nil
}

func (d *DataAvailabilityProvider) StoreBlobStatus(ctx context.Context, requestId []byte, batches []common.Hash) error {
	blobStatusReply, err := d.client.GetBlobStatus(ctx, requestId)
	if err != nil {
		fmt.Printf("error getting blob status: %v\n", err)
		return err
	}

	// Process and store blob status in storage
	currStatus := blobStatusReply.GetStatus()
	if currStatus == disperser_rpc.BlobStatus_CONFIRMED || currStatus == disperser_rpc.BlobStatus_FINALIZED {
		blobInfo := blobStatusReply.GetInfo()
		blobVerificationProof := blobInfo.GetBlobVerificationProof()

		// Store blob information inside in-memory DA storage
		for idx, hash := range batches {
			err := d.state.Add(hash, blobVerificationProof)
			if err != nil {
				fmt.Printf("error adding blob into storage: %v\n", err)
				// Should not come here, but we will panic the mock node if indexing fails
				panic(err)
			}
			err = d.state.AddIndex(hash, idx)
			if err != nil {
				fmt.Printf("error adding batch index into storage: %v\n", err)
				// Should not come here, but we will panic the mock node if indexing fails
				panic(err)
			}
		}
		return nil
	}
	return fmt.Errorf("failed to store blob in DA storage, blob is not confirmed or finalized")
}

// GetBatchL2Data returns the data from the EigenDA layer operators. It checks the DA storage to get the
// requestID used when submitting the batch data to the DA.
func (d *DataAvailabilityProvider) GetBatchL2Data(ctx context.Context, hash common.Hash) ([]byte, error) {
	blobInfo, err := d.state.Get(hash)
	if err != nil {
		fmt.Println("failed to get blob info from DA storage")
		return nil, err
	}
	reply, err := d.client.RetrieveBlob(ctx, blobInfo.BatchHeaderHash, blobInfo.BlobIndex)
	if err != nil {
		fmt.Printf("failed to retrieve blob: %v\n", err)
		return nil, err
	}
	idx, err := d.state.GetIndex(hash)
	if err != nil {
		fmt.Println("failed to get blob index from DA storage")
		return nil, err
	}
	data := reply.GetData()
	batchesData := DecodeSequence(data)
	return batchesData[idx], nil
}

// EncodeSequence is the helper function to encode sequence data into 1D byte array. The
// encoding scheme is ensured to be lossless.
//
// The first n+1 8-bytes of the blob contains the metadata of the batches data.
// The first 8-bytes stores the size of the sequence, and the next 8-bytes will store the
// byte array length of every batch data.
func EncodeSequence(batchesData [][]byte) []byte {
	sequence := []byte{}
	metadata := []byte{}
	n := uint64(len(batchesData))
	bn := make([]byte, 8)
	binary.BigEndian.PutUint64(bn, n)
	metadata = append(metadata, bn...)

	for _, seq := range batchesData {
		// Add batch data to byte array
		sequence = append(sequence, seq...)

		// Add batch metadata to byte array
		n := uint64(len(seq))
		bn := make([]byte, 8)
		binary.BigEndian.PutUint64(bn, n)
		metadata = append(metadata, bn...)
	}
	sequence = append(metadata, sequence...)
	return sequence
}

// DecodeSequence is the helper function to decode 1D byte array into sequence data. The
// encoding scheme is ensured to be lossless.
//
// When decoding the blob data, the first n+1 8-bytes of the blob contains the metadata of
// the batches data.
// The first 8-bytes stores the size of the sequence, and the next 8-bytes will store the
// byte array length of every batch data.
func DecodeSequence(blobData []byte) [][]byte {
	bn := blobData[:8]
	n := binary.BigEndian.Uint64(bn)
	metadata := blobData[8 : 8*(n+1)]
	sequence := blobData[8*(n+1):]

	batchesData := [][]byte{}
	idx := uint64(0)
	for i := uint64(0); i < n; i++ {
		// Get batch data byte array length
		bn := metadata[8*i : 8*i+8]
		n := binary.BigEndian.Uint64(bn)
		batchesData = append(batchesData, sequence[idx:idx+n])
		idx += n
	}
	return batchesData
}
