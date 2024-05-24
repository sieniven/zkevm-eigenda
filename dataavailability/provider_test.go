package dataavailability

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/Layr-Labs/eigenda/encoding/utils/codec"
	"github.com/sieniven/zkevm-eigenda/config/types"
	"github.com/stretchr/testify/assert"
)

func TestDisperseBlobWithStringDataUsingProvider(t *testing.T) {
	cfg := Config{
		Hostname:          "disperser-holesky.eigenda.xyz",
		Port:              "443",
		Timeout:           types.NewDuration(30 * time.Second),
		UseSecureGrpcFlag: true,
	}
	provider := NewDataProvider(cfg)

	// Generate mock string batch data
	data := []byte("hihihihihihihihihihihihihihihihihihi")
	data = codec.ConvertByPaddingEmptyByte(data)

	// Generate mock string sequence
	mockBatches := [][]byte{}
	for i := 0; i < 10; i++ {
		mockBatches = append(mockBatches, data)
	}

	// send mock sequence with provider
	blobInfo, err := provider.PostSequence(context.Background(), mockBatches)
	assert.NoError(t, err)
	assert.NotZero(t, blobInfo.BlobIndex)
	assert.NotNil(t, blobInfo.BatchHeaderHash)
	assert.NotEmpty(t, blobInfo.BatchHeaderHash)
	assert.NotNil(t, blobInfo.BatchRoot)
	assert.NotEmpty(t, blobInfo.BatchRoot)
	assert.NotZero(t, blobInfo.ReferenceBlockNumber)
}

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
	data = codec.ConvertByPaddingEmptyByte(data)

	// Generate mock string sequence
	mockBatches := [][]byte{}
	for i := 0; i < 10; i++ {
		mockBatches = append(mockBatches, data)
	}

	// send mock sequence with provider
	blobInfo, err := provider.PostSequence(context.Background(), mockBatches)
	assert.NoError(t, err)
	assert.NotZero(t, blobInfo.BlobIndex)
	assert.NotNil(t, blobInfo.BatchHeaderHash)
	assert.NotEmpty(t, blobInfo.BatchHeaderHash)
	assert.NotNil(t, blobInfo.BatchRoot)
	assert.NotEmpty(t, blobInfo.BatchRoot)
	assert.NotZero(t, blobInfo.ReferenceBlockNumber)
}
