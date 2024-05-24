package dataavailability

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Layr-Labs/eigenda/encoding/utils/codec"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecodeSequenceToAndFromStringBlob(t *testing.T) {
	mock_string_data := "hihihihihihihihihihihihihihihihihihi"
	data := []byte(mock_string_data)

	// Generate mock sequence data
	mockSeqData := [][]byte{}
	for i := 0; i < 10; i++ {
		mockSeqData = append(mockSeqData, data)
	}
	blob := EncodeSequence(mockSeqData)

	// Decode blob
	decodedBlob := DecodeSequence(blob)

	// Assert decoded sequence length is correct
	n := len(decodedBlob)
	assert.Equal(t, 10, n)

	// Assert decoded sequence data is correct
	for i := 0; i < n; i++ {
		data_decoded := string(decodedBlob[i])
		assert.Equal(t, data_decoded, mock_string_data)
	}
}

func TestEncodeDecodeSequenceToAndFromRandomBlob(t *testing.T) {
	// Define Different DataSizes
	dataSize := []int{100000, 200000, 1000, 80, 30000}

	// Generate mock sequence data
	mockSeqData := [][]byte{}
	for i := 0; i < 10; i++ {
		// Disperse Blob with different DataSizes
		rand.Seed(time.Now().UnixNano())
		data := make([]byte, dataSize[rand.Intn(len(dataSize))])
		_, err := rand.Read(data)
		assert.NoError(t, err)

		data = codec.ConvertByPaddingEmptyByte(data)
		mockSeqData = append(mockSeqData, data)
	}
	blob := EncodeSequence(mockSeqData)

	// Decode blob
	decodedBlob := DecodeSequence(blob)

	// Assert decoded sequence length is correct
	n := len(decodedBlob)
	assert.Equal(t, 10, n)

	// Assert decoded sequence data is correct
	for i := 0; i < n; i++ {
		data_decoded := decodedBlob[i]
		assert.Equal(t, data_decoded, mockSeqData[i])
	}
}
