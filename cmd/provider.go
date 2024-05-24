package main

import (
	"context"
	"fmt"

	"github.com/Layr-Labs/eigenda/encoding/utils/codec"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/urfave/cli/v2"
)

func testDAProvider(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return err
	}
	setupLog(c.Log)

	provider := dataavailability.NewDataProvider(c.EigenDAClient)

	// Generate mock string batch data
	data := []byte("hihihihihihihihihihihihihihihihihihi")
	data = codec.ConvertByPaddingEmptyByte(data)

	// Generate mock string sequence
	mockBatches := [][]byte{}
	for i := 0; i < 10; i++ {
		mockBatches = append(mockBatches, data)
	}

	// Test PostSequence pipeline
	ctx := context.Background()
	info, err := provider.PostSequence(ctx, mockBatches)
	if err != nil {
		panic(err)
	}
	fmt.Println("blob info: ", info)

	// Test GetSequence pipeline
	return nil
}
