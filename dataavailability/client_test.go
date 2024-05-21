package dataavailability

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	disperser_rpc "github.com/Layr-Labs/eigenda/api/grpc/disperser"
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

func TestClientDisperseBlockAndGetBlobConfirmation(t *testing.T) {
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
	var status disperser_rpc.BlobStatus
	timer := time.Now()
	for {
		blobStatusReply, err := client.GetBlobStatus(ctx, idBytes)
		assert.NoError(t, err)
		assert.NotNil(t, blobStatusReply)

		// Log blob status
		currStatus := blobStatusReply.GetStatus()
		if currStatus != status {
			prevStatusStr := disperser_rpc.BlobStatus_name[int32(status)]
			currStatusStr := disperser_rpc.BlobStatus_name[int32(currStatus)]
			elapsed := time.Since(timer)
			fmt.Println("---- Blob state ----")
			fmt.Printf("Change of state from current blob status %v to new blob status %v took: %s\n", prevStatusStr, currStatusStr, elapsed)

			// Log blob information
			message := blobStatusReply.String()
			fmt.Println("---- Blob information ----")
			fmt.Print(message + "\n\n")

			// Reset timer and status
			timer = time.Now()
			status = currStatus
			time.Sleep(10 * time.Second)
		}
	}
}
