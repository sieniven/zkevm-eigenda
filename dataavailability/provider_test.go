package dataavailability

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sieniven/zkevm-eigenda/config/types"
	"github.com/stretchr/testify/assert"
)

// Set longer timeout flag for test case
func TestDisperseBlobWithStringDataUsingProvider(t *testing.T) {
	cfg := Config{
		Hostname:          "disperser-holesky.eigenda.xyz",
		Port:              "443",
		Timeout:           types.NewDuration(30 * time.Second),
		UseSecureGrpcFlag: true,
	}
	provider := NewDataProvider(cfg)

	// Generate mock string batch data
	stringData := "hihihihihihihihihihihihihihihihihihi"
	data := []byte(stringData)

	// Generate mock string sequence
	mockBatches := [][]byte{}
	for i := 0; i < 10; i++ {
		mockBatches = append(mockBatches, data)
	}

	// send mock sequence with provider
	msg, err := provider.PostSequence(context.Background(), mockBatches)
	assert.NoError(t, err)
	proof, err := TryFromDataAvailabilityMessage(msg)
	assert.NoError(t, err)
	assert.NotZero(t, proof.BatchId)
	assert.NotZero(t, proof.BlobIndex)
	assert.NotNil(t, proof.BatchMetadata.BatchHeader.BlobHeadersRoot.Bytes())
	assert.NotEmpty(t, proof.BatchMetadata.BatchHeader.BlobHeadersRoot.Bytes())
	assert.NotNil(t, proof.BatchMetadata.BatchHeader.QuorumNumbers)
	assert.NotEmpty(t, proof.BatchMetadata.BatchHeader.QuorumNumbers)
	assert.NotNil(t, proof.BatchMetadata.BatchHeader.SignedStakeForQuorums)
	assert.NotEmpty(t, proof.BatchMetadata.BatchHeader.SignedStakeForQuorums)
	assert.NotZero(t, proof.BatchMetadata.BatchHeader.ReferenceBlockNumber)
	assert.NotNil(t, proof.BatchMetadata.SignatoryRecordHash.Bytes())
	assert.NotEmpty(t, proof.BatchMetadata.SignatoryRecordHash.Bytes())
	assert.NotZero(t, proof.BatchMetadata.ConfirmationBlockNumber)
	assert.NotNil(t, proof.InclusionProof)
	assert.NotEmpty(t, proof.InclusionProof)
	assert.NotNil(t, proof.QuorumIndices)
	assert.NotEmpty(t, proof.QuorumIndices)

	// Retrieve sequence with provider
	batchesData, err := provider.GetSequence(context.Background(), []common.Hash{}, msg)

	// Validate retrieved data
	assert.NoError(t, err)
	assert.Equal(t, 10, len(batchesData))
	for _, batchData := range batchesData {
		assert.Equal(t, stringData, string(batchData))
	}
}

// Set longer timeout flag for test case
func TestDisperseBlobWithRandomDataUsingProvider(t *testing.T) {
	cfg := Config{
		Hostname:          "disperser-holesky.eigenda.xyz",
		Port:              "443",
		Timeout:           types.NewDuration(30 * time.Second),
		UseSecureGrpcFlag: true,
	}
	provider := NewDataProvider(cfg)

	// Define Different DataSizes
	dataSize := []int{100000, 200000, 1000, 80, 30000}

	// Disperse Blob with different DataSizes
	rand.Seed(time.Now().UnixNano())
	data := make([]byte, dataSize[rand.Intn(len(dataSize))])
	_, err := rand.Read(data)
	assert.NoError(t, err)

	// Generate mock string sequence
	mockBatches := [][]byte{}
	for i := 0; i < 10; i++ {
		mockBatches = append(mockBatches, data)
	}

	// send mock sequence with provider
	msg, err := provider.PostSequence(context.Background(), mockBatches)
	assert.NoError(t, err)
	proof, err := TryFromDataAvailabilityMessage(msg)
	assert.NoError(t, err)
	assert.NotZero(t, proof.BatchId)
	assert.NotZero(t, proof.BlobIndex)
	assert.NotNil(t, proof.BatchMetadata.BatchHeader.BlobHeadersRoot.Bytes())
	assert.NotEmpty(t, proof.BatchMetadata.BatchHeader.BlobHeadersRoot.Bytes())
	assert.NotNil(t, proof.BatchMetadata.BatchHeader.QuorumNumbers)
	assert.NotEmpty(t, proof.BatchMetadata.BatchHeader.QuorumNumbers)
	assert.NotNil(t, proof.BatchMetadata.BatchHeader.SignedStakeForQuorums)
	assert.NotEmpty(t, proof.BatchMetadata.BatchHeader.SignedStakeForQuorums)
	assert.NotZero(t, proof.BatchMetadata.BatchHeader.ReferenceBlockNumber)
	assert.NotNil(t, proof.BatchMetadata.SignatoryRecordHash.Bytes())
	assert.NotEmpty(t, proof.BatchMetadata.SignatoryRecordHash.Bytes())
	assert.NotZero(t, proof.BatchMetadata.ConfirmationBlockNumber)
	assert.NotNil(t, proof.InclusionProof)
	assert.NotEmpty(t, proof.InclusionProof)
	assert.NotNil(t, proof.QuorumIndices)
	assert.NotEmpty(t, proof.QuorumIndices)

	// Retrieve sequence with provider
	batchesData, err := provider.GetSequence(context.Background(), []common.Hash{}, msg)

	// Validate retrieved data
	assert.NoError(t, err)
	assert.Equal(t, 10, len(batchesData))
	for idx, batchData := range batchesData {
		assert.Equal(t, mockBatches[idx], batchData)
	}
}
