package dataavailability

import (
	"context"
	"encoding/binary"
	"fmt"
	"sync"

	"github.com/Layr-Labs/eigenda/clients"
	"github.com/Layr-Labs/eigenda/encoding/utils/codec"
	"github.com/ethereum/go-ethereum/common"
)

// DataAvailabilityProvider is the EigenDA backend manager that holds the DA implementation.
// It contains the implementation of SequenceSender and SequenceRetriever of the zkevm's
// dataavailability package.
type DataAvailabilityProvider struct {
	cfg       clients.Config
	state     DAStorage
	disperser clients.DisperserClient
	retriever clients.RetrievalClient
}

// Factory method for a new DataAvailibilityProvider instance
func NewDataProvider(cfg Config) *DataAvailabilityProvider {
	// Initialize in-memory DA storage
	s := DAStorage{
		inner: map[common.Hash][]byte{},
		mutex: &sync.RWMutex{},
	}
	c := clients.Config{
		Hostname:          cfg.Hostname,
		Port:              cfg.Port,
		Timeout:           cfg.Timeout.Duration,
		UseSecureGrpcFlag: cfg.UseSecureGrpcFlag,
	}
	signer := MockBlobRequestSigner{}

	p := &DataAvailabilityProvider{
		cfg:       c,
		state:     s,
		disperser: clients.NewDisperserClient(&c, signer),
		retriever: clients.NewRetrievalClient(),
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
	_, idBytes, err := d.disperser.DisperseBlob(ctx, blobData, []uint8{})
	if err != nil {
		fmt.Println("failed to send blob to EigenDA disperser")
		return []byte{}, nil
	}

	fmt.Println("sent blob to EigenDA disperser")
	return idBytes, nil
}

func (d *DataAvailabilityProvider) GetSequence(ctx context.Context, batchHashes []common.Hash, dataAvailabilityMessage []byte) ([][]byte, error) {
	// Get
}

// GetBatchL2Data returns the data from the EigenDA layer operators. It checks the DA storage to get the
// requestID used when submitting the batch data to the DA.
func (d *DataAvailabilityProvider) GetBatchL2Data(hash common.Hash) ([]byte, error) {
	id, err := d.state.Get(hash)
	if err != nil {
		fmt.Println("failed to get blob requestID from DA storage")
		return nil, err
	}
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
