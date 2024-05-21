package dataavailability

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/Layr-Labs/eigenda/disperser"
	"github.com/Layr-Labs/eigenda/encoding/utils/codec"
	"github.com/sieniven/zkevm-eigenda/config/types"
	daTypes "github.com/sieniven/zkevm-eigenda/dataavailability/types"
	"github.com/stretchr/testify/assert"
)

func TestClientDisperseBlob(t *testing.T) {
	cfg := Config{
		Hostname:          "disperser-holesky.eigenda.xyz",
		Port:              "443",
		Timeout:           types.NewDuration(30 * time.Second),
		UseSecureGrpcFlag: true,
	}
	signer := daTypes.MockBlobRequestSigner{}
	client := NewDisperserClient(&cfg, signer)

	// Define Different DataSizes
	dataSize := []int{100000, 200000, 1000, 80, 30000}

	// Disperse Blob with different DataSizes
	rand.Seed(time.Now().UnixNano())
	data := make([]byte, dataSize[rand.Intn(len(dataSize))])
	_, err := rand.Read(data)
	assert.NoError(t, err)

	data = codec.ConvertByPaddingEmptyByte(data)

	// Send blob
	blobStatus, idBytes, err := client.DisperseBlob(context.Background(), data, []uint8{})
	assert.NoError(t, err)
	assert.NotNil(t, blobStatus)
	assert.Equal(t, *blobStatus, disperser.Processing)
	assert.True(t, len(idBytes) > 0)
	id := string(idBytes)
	fmt.Println("id: ", id)
}
