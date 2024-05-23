package dataavailability

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/Layr-Labs/eigenda/clients"
	"github.com/Layr-Labs/eigenda/disperser"
	"github.com/Layr-Labs/eigenda/encoding/utils/codec"
	"github.com/stretchr/testify/assert"
)

func TestDisperserClientDisperseBlobWithStringData(t *testing.T) {
	cfg := clients.Config{
		Hostname:          "disperser-holesky.eigenda.xyz",
		Port:              "443",
		Timeout:           time.Duration(30 * time.Second),
		UseSecureGrpcFlag: true,
	}
	signer := MockBlobRequestSigner{}
	client := clients.NewDisperserClient(&cfg, signer)

	data := []byte("hihihihihihihihihihihihihihihihihihi")
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

func TestDisperserClientDisperseBlobWithRandomData(t *testing.T) {
	cfg := clients.Config{
		Hostname:          "disperser-holesky.eigenda.xyz",
		Port:              "443",
		Timeout:           time.Duration(30 * time.Second),
		UseSecureGrpcFlag: true,
	}
	signer := MockBlobRequestSigner{}
	client := clients.NewDisperserClient(&cfg, signer)

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
