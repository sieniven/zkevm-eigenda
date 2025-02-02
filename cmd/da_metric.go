package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	disperser_rpc "github.com/Layr-Labs/eigenda/api/grpc/disperser"
	"github.com/Layr-Labs/eigenda/encoding/utils/codec"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability/eigenda"
	"github.com/urfave/cli/v2"
)

func getEigenDAMetrics(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return err
	}
	setupLog(c.Log)

	signer := eigenda.MockBlobRequestSigner{}
	client := eigenda.NewDisperserClient(&c.DataAvailability, signer)

	// Generate mock string batch data
	stringData := "hihihihihihihihihihihihihihihihihihi"
	data := []byte(stringData)
	data = codec.ConvertByPaddingEmptyByte(data)
	ctx := context.Background()

	// Send blob
	_, idBytes, err := client.DisperseBlob(ctx, data, []uint8{})
	if err != nil {
		panic(err)
	}
	id := base64.StdEncoding.EncodeToString(idBytes)
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
			time.Sleep(c.DataAvailability.RetrieveBlobStatusPeriod.Duration)

			if status == disperser_rpc.BlobStatus_CONFIRMED {
				// Test retrieve blob pipeline
				blobStatusReply, err := client.GetBlobStatus(ctx, idBytes)
				if err != nil {
					panic(err)
				}
				if blobStatusReply == nil {
					panic(fmt.Errorf("empty blob status reply returned"))
				}

				info := blobStatusReply.GetInfo()
				blobData, err := eigenda.GetBlobData(info)
				if err != nil {
					panic(err)
				}
				batchHeaderHash, err := blobData.BlobVerificationProof.BatchMetadata.GetBatchHeaderHash()
				if err != nil {
					panic(err)
				}

				reply, err := client.RetrieveBlob(ctx, batchHeaderHash, blobData.BlobVerificationProof.BlobIndex)
				if err != nil {
					panic(err)
				}

				retrievedData := reply.GetData()
				retrievedData = codec.RemoveEmptyByteFromPaddedBytes(retrievedData)
				if string(retrievedData) != stringData {
					panic(fmt.Errorf("retrieved data does not equal to initial data"))
				}
				fmt.Println("decoded batch data: ", retrievedData)
			}
		}

		if status == disperser_rpc.BlobStatus_FINALIZED {
			// Break once blob status is finalized
			break
		}
	}
	return nil
}
