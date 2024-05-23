package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	disperser_rpc "github.com/Layr-Labs/eigenda/api/grpc/disperser"
	"github.com/Layr-Labs/eigenda/clients"
	"github.com/Layr-Labs/eigenda/encoding/utils/codec"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/urfave/cli/v2"
)

func getEigenDAMetrics(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return err
	}

	cfg := clients.Config{
		Hostname:          c.EigenDAClient.Hostname,
		Port:              c.EigenDAClient.Port,
		Timeout:           c.EigenDAClient.Timeout.Duration,
		UseSecureGrpcFlag: c.EigenDAClient.UseSecureGrpcFlag,
	}
	signer := dataavailability.MockBlobRequestSigner{}
	client := clients.NewDisperserClient(&cfg, signer)

	// Define Different DataSizes
	dataSize := []int{100000, 200000, 1000, 80, 30000}

	// Disperse Blob with different DataSizes
	rand.Seed(time.Now().UnixNano())
	data := make([]byte, dataSize[rand.Intn(len(dataSize))])
	_, err = rand.Read(data)
	if err != nil {
		panic(err)
	}

	data = codec.ConvertByPaddingEmptyByte(data)
	ctx := context.Background()

	// Send blob
	_, idBytes, err := client.DisperseBlob(ctx, data, []uint8{})
	if err != nil {
		panic(err)
	}
	id := string(idBytes)
	fmt.Println("id: ", id)

	// Monitor blob status
	var status disperser_rpc.BlobStatus
	timer := time.Now()
	for {
		blobStatusReply, err := client.GetBlobStatus(ctx, idBytes)
		if err != nil {
			panic(err)
		}

		// Log blob status
		currStatus := blobStatusReply.GetStatus()
		if currStatus != status {
			prevStatusStr, ok := disperser_rpc.BlobStatus_name[int32(status)]
			if !ok {
				fmt.Println("Caught an unexpected status, prev status: ", status)
				prevStatusStr = string(int32(status))
			}
			currStatusStr, ok := disperser_rpc.BlobStatus_name[int32(currStatus)]
			if !ok {
				fmt.Println("Caught an unexpected status, curr status: ", currStatus)
				currStatusStr = string(int32(currStatus))
			}
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
	return nil
}
