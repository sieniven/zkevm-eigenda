package main

import (
	"encoding/base64"
	"fmt"

	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/sieniven/zkevm-eigenda/dataavailability/eigenda"
	"github.com/urfave/cli/v2"
)

func retrieve(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return nil
	}
	setupLog(c.Log)

	p := eigenda.NewDataAvailabilityProvider(c.DataAvailability)
	da := dataavailability.New(c.DataAvailability, p)

	// Get EigenDA blob information
	requestId := cliCtx.String(config.FlagRequestID)
	id, err := base64.StdEncoding.DecodeString(requestId)
	if err != nil {
		panic(err)
	}
	batchesData, err := da.GetBatchL2DataFromRequestId(cliCtx.Context, id)
	if err != nil {
		fmt.Println("failed to get batch data from req id: ", err)
		return err
	}

	fmt.Println("Retrieved batches data of length ", len(batchesData))

	return nil
}
