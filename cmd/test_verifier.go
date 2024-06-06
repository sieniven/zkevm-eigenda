package main

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/sieniven/zkevm-eigenda/dataavailability/eigenda"
	"github.com/urfave/cli/v2"
)

func testVerifier(cliCtx *cli.Context) error {
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
	p := eigenda.NewDataAvailabilityProvider(c.DataAvailability)
	da := dataavailability.New(c.DataAvailability, p)
	etherMan.SetDataProvider(da)
	_, privKey, err := etherMan.LoadAuthFromKeyStore(c.Key.Path, c.Key.Password)
	if err != nil {
		panic(err)
	}
	if privKey == nil { //nolint:staticcheck
		panic(errors.New("private key not found"))
	}

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

	for _, batchData := range batchesData {
		for _, b := range batchData {
			if b != byte(10) {
				panic(fmt.Errorf("invalid batch data"))
			}
		}
		fmt.Println("Valid batch data :)")
	}

	fmt.Println("Provider posted and retrieved valid batches data from EigenDA layer")

	protocolAddr, err := etherMan.EigendaVerifier.GetDataAvailabilityProtocol(&bind.CallOpts{Pending: false})
	if err != nil {
		panic(err)
	}
	fmt.Println("Protocol address: ", protocolAddr.Hex())

	// Test on-chain verification
	fmt.Println("Verifying on-chain...")
	daMessage, err := da.GetDataAvailabilityMessageFromRequestId(cliCtx.Context, id)
	if err != nil {
		panic(err)
	}
	err = etherMan.VerifyDataAvailabilityMessage(daMessage)
	if err != nil {
		panic(err)
	}
	fmt.Println("Verified!")

	return nil
}
