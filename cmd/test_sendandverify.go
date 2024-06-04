package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/sieniven/zkevm-eigenda/dataavailability/eigenda"
	"github.com/sieniven/zkevm-eigenda/etherman/types"
	batchTypes "github.com/sieniven/zkevm-eigenda/sequencesender/types"
	"github.com/urfave/cli/v2"
)

func testSendAndVerify(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return err
	}
	setupLog(c.Log)

	// Initialize ether manager instance
	etherMan, err := newEtherman(*c)
	if err != nil {
		panic(err)
	}

	// Create new data avaiability manager
	p := eigenda.NewDataProvider(c.DataAvailability)
	da := dataavailability.New(c.DataAvailability, p)
	etherMan.SetDataProvider(da)
	_, privKey, err := etherMan.LoadAuthFromKeyStore(c.Key.Path, c.Key.Password)
	if err != nil {
		panic(err)
	}
	if privKey == nil { //nolint:staticcheck
		panic(errors.New("private key not found"))
	}

	// Generate mock batch data
	// Generate mock batch data for max configured size
	data := make([]byte, c.SequenceSender.MaxBatchBytesSize)
	for i := uint64(0); i < c.SequenceSender.MaxBatchBytesSize; i++ {
		data[i] = byte(10)
	}

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

	daMessage, err := da.PostSequence(cliCtx.Context, sequence)
	if err != nil {
		panic(err)
	}
	fmt.Println("blob info: ", daMessage)

	// Retrieve sequence
	batchesData, err := da.GetBatchL2Data(batchNums, []common.Hash{}, daMessage)
	if err != nil {
		panic(err)
	}

	for _, batchData := range batchesData {
		for _, b := range batchData {
			if b != byte(10) {
				panic(fmt.Errorf("invalid batch data"))
			}
		}
		fmt.Println("Valid batch data :)")
	}

	fmt.Println("Provider posted and retrieved valid batches data from EigenDA layer")

	// Test on-chain verification
	err = etherMan.VerifyDataAvailabilityMessage(daMessage)
	if err != nil {
		panic(err)
	}

	return nil
}
