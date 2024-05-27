package main

import (
	"fmt"

	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/urfave/cli/v2"
)

func retrieve(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return nil
	}
	setupLog(c.Log)

	da := dataavailability.New(c.EigenDAClient)

	// Get EigenDA blob information
	requestId := cliCtx.String(config.RequestID)
	id := []byte(requestId)
	batchesData, err := da.GetBatchL2DataFromRequestId(cliCtx.Context, id)
	if err != nil {
		fmt.Println("failed to get batch data from req id: ", err)
		return err
	}

	for idx, batchData := range batchesData {
		fmt.Printf("Batch %v data: %v\n", idx, string(batchData))
	}

	return nil
}
