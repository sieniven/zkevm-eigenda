package main

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/sieniven/zkevm-eigenda/etherman/types"
	batchTypes "github.com/sieniven/zkevm-eigenda/sequencesender/types"
	"github.com/urfave/cli/v2"
)

func testDAProvider(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return err
	}
	setupLog(c.Log)

	da := dataavailability.New(c.EigenDAClient)

	// Generate mock batch data
	stringData := "hihihihihihihihihihihihihihihihihihi"
	data := []byte(stringData)

	batches := []batchTypes.Batch{}
	batchNums := []uint64{}
	for i := 0; i < 10; i++ {
		batch := batchTypes.Batch{
			BatchNumber: uint64(i),
			BatchL2Data: data,
			Timestamp:   time.Now(),
		}
		batches = append(batches, batch)
		batchNums = append(batchNums, uint64(i))
	}

	sequence := []types.Sequence{}
	for _, batch := range batches {
		seq := types.Sequence{
			BatchL2Data: batch.BatchL2Data,
			BatchNumber: batch.BatchNumber,
		}
		sequence = append(sequence, seq)
	}

	info, err := da.PostSequence(cliCtx.Context, sequence)
	if err != nil {
		panic(err)
	}
	fmt.Println("blob info: ", info)

	// Retrieve sequence
	batchesData, err := da.GetBatchL2Data(batchNums, []common.Hash{}, info)
	if err != nil {
		panic(err)
	}

	for _, batch := range batchesData {
		if stringData == string(batch) {
			fmt.Println("valid batch data :)")
		} else {
			panic(fmt.Errorf("invalid batch data"))
		}
	}

	// Generate mock string sequence
	return nil
}
