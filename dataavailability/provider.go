package dataavailability

import (
	"context"
	"fmt"
	"sync"
	"time"

	disperser_rpc "github.com/Layr-Labs/eigenda/api/grpc/disperser"
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
		inner: map[common.Hash]BlobInfo{},
		mutex: &sync.RWMutex{},
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

func (d *DataAvailabilityProvider) PostSequence(ctx context.Context, batchesData [][]byte) (BlobInfo, error) {
	// Send blob to EigenDA disperser
	blobData := EncodeSequence(batchesData)
	_, idBytes, err := d.client.DisperseBlob(ctx, blobData, []uint8{})
	if err != nil {
		fmt.Println("failed to send blob to EigenDA disperser: ", err)
		return BlobInfo{}, nil
	}
	fmt.Println("sent blob to EigenDA disperser, request id: ", string(idBytes))

	var blobStatusReply *disperser_rpc.BlobStatusReply
	for {
		blobStatusReply, err = d.client.GetBlobStatus(ctx, idBytes)
		if err != nil {
			fmt.Printf("error getting blob status: %v\n", err)
			return BlobInfo{}, err
		}

		// Get blob status
		currStatus := blobStatusReply.GetStatus()
		if currStatus == disperser_rpc.BlobStatus_CONFIRMED || currStatus == disperser_rpc.BlobStatus_FINALIZED {
			break
		}

		// Wait period before retrieving blob status
		time.Sleep(d.cfg.RetrieveBlobStatusPeriod.Duration)
	}

	if blobStatusReply == nil {
		err = fmt.Errorf("empty blob status reply returned")
		return BlobInfo{}, err
	}

	info := blobStatusReply.GetInfo()
	blob := info.GetBlobVerificationProof()
	blobInfo := BlobInfo{
		BlobIndex:            blob.BlobIndex,
		BatchHeaderHash:      blob.BatchMetadata.BatchHeaderHash,
		BatchRoot:            blob.BatchMetadata.BatchHeader.BatchRoot,
		ReferenceBlockNumber: uint(blob.BatchMetadata.ConfirmationBlockNumber),
	}

	return blobInfo, nil
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
	data, _ := DecodeSequence(reply.GetData())
	batchesData = append(batchesData, data...)

	return batchesData, nil
}

func (d *DataAvailabilityProvider) StoreBlobStatus(ctx context.Context, batchHash common.Hash, blobInfo BlobInfo) error {
	// Store blob information inside in-memory DA storage
	err := d.state.Add(batchHash, blobInfo)
	if err != nil {
		fmt.Printf("error adding blob into storage: %v\n", err)
		// Should not come here, but we will panic the mock node if indexing fails
		panic(err)
	}

	return nil
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
	data := reply.GetData()
	batchesData, batchesHash := DecodeSequence(data)

	// Get batch data from batch hash
	for idx, h := range batchesHash {
		if h == hash {
			return batchesData[idx], nil
		}
	}
	return nil, fmt.Errorf("failed to get batch data from hash, corrupted DA storage")
}

// Get blob information from request ID
func (d *DataAvailabilityProvider) GetBlobInformationFromId(ctx context.Context, requestId []byte) (BlobInfo, error) {
	blobStatusReply, err := d.client.GetBlobStatus(ctx, requestId)
	if err != nil {
		fmt.Printf("error getting blob status: %v\n", err)
		return BlobInfo{}, err
	}

	// Get blob status
	status := blobStatusReply.GetStatus()
	confirmedFlag := status == disperser_rpc.BlobStatus_CONFIRMED || status == disperser_rpc.BlobStatus_FINALIZED

	if confirmedFlag {
		info := blobStatusReply.GetInfo()
		blob := info.GetBlobVerificationProof()
		blobInfo := BlobInfo{
			BlobIndex:            blob.BlobIndex,
			BatchHeaderHash:      blob.BatchMetadata.BatchHeaderHash,
			BatchRoot:            blob.BatchMetadata.BatchHeader.BatchRoot,
			ReferenceBlockNumber: uint(blob.BatchMetadata.ConfirmationBlockNumber),
		}
		return blobInfo, nil
	} else {
		return BlobInfo{}, fmt.Errorf("EigenDA blob not confirmed, unable to retrieve blob information")
	}
}
