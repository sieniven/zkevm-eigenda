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
	"github.com/stretchr/testify/assert"
)

func TestClientDisperseBlob(t *testing.T) {
	cfg := Config{
		Hostname:          "disperser-holesky.eigenda.xyz",
		Port:              "443",
		Timeout:           types.NewDuration(30 * time.Second),
		UseSecureGrpcFlag: true,
	}
	signer := MockBlobRequestSigner{}
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

func TestClientDisperseBlockAndGetBlobConfirmation(t *testing.T) {
	cfg := Config{
		Hostname:          "disperser-holesky.eigenda.xyz",
		Port:              "443",
		Timeout:           types.NewDuration(30 * time.Second),
		UseSecureGrpcFlag: true,
	}
	signer := MockBlobRequestSigner{}
	client := NewDisperserClient(&cfg, signer)

	// Define Different DataSizes
	dataSize := []int{100000, 200000, 1000, 80, 30000}

	// Disperse Blob with different DataSizes
	rand.Seed(time.Now().UnixNano())
	data := make([]byte, dataSize[rand.Intn(len(dataSize))])
	_, err := rand.Read(data)
	assert.NoError(t, err)

	data = codec.ConvertByPaddingEmptyByte(data)
	ctx := context.Background()

	// Send blob
	blobStatus, idBytes, err := client.DisperseBlob(ctx, data, []uint8{})
	assert.NoError(t, err)
	assert.NotNil(t, blobStatus)
	assert.Equal(t, *blobStatus, disperser.Processing)
	assert.True(t, len(idBytes) > 0)
	id := string(idBytes)
	fmt.Println("id: ", id)

	// Monitor blob status
	for {
		time.Sleep(10 * time.Second)
		fmt.Println("trying to get dispersed blob status...")

		blobStatusReply, err := client.GetBlobStatus(ctx, idBytes)
		assert.NoError(t, err)
		assert.NotNil(t, blobStatusReply)
		message := blobStatusReply.String()
		fmt.Println(message)

		switch status := blobStatusReply.GetStatus(); status {
		case 0:
			fmt.Println("Blob status: UNKNOWN")
		case 1:
			fmt.Println("Blob status: PROCESSING")
		case 2:
			fmt.Println("Blob status: CONFIRMED")
		case 3:
			fmt.Println("Blob status: FAILED")
		case 4:
			fmt.Println("Blob status: FINALIZED")
		case 5:
			fmt.Println("Blob status: INSUFFICIENT SIGNATURES")
		case 6:
			fmt.Println("Blobk status: DISPERSING")
		default:
			t.Fail()
		}
	}
}
