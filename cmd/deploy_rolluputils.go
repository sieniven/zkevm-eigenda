package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/sieniven/zkevm-eigenda/dataavailability/eigenda"
	"github.com/sieniven/zkevm-eigenda/ethtxmanager"
	"github.com/urfave/cli/v2"
)

func deployLibraries(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return nil
	}
	setupLog((c.Log))

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

	admin := cliCtx.String(config.FlagAdmin)
	adminAddr := common.HexToAddress(admin)
	fmt.Println("Deploying with admin address: ", adminAddr)
	fmt.Printf("from pk %s, from sender %s\n", crypto.PubkeyToAddress(privKey.PublicKey), adminAddr.String()) //nolint:staticcheck

	// Get EigenDARollupUtils library compiled bytecode
	content, err := os.ReadFile("/Users/nivensie/zkevm-eigenda/etherman/smartcontracts/bin/eigendarolluputils.bin")
	if err != nil {
		panic(err)
	}
	bytecode, err := hex.DecodeString(string(content))
	if err != nil {
		panic(err)
	}

	// Start eth-tx-manager service
	etm := ethtxmanager.New(c.EthTxManager, etherMan)
	go etm.Start()

	// Send library deployment tx
	id := "deploy"
	owner := "deployment"
	value := big.NewInt(0)
	err = etm.Add(cliCtx.Context, owner, id, adminAddr, nil, value, bytecode, c.SequenceSender.GasOffset)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sent signed tx to node")

	flag := true
	for flag {
		etm.ProcessPendingMonitoredTxs(cliCtx.Context, owner, func(result ethtxmanager.MonitoredTxResult) {
			if result.Status == ethtxmanager.MonitoredTxStatusFailed || result.Status == ethtxmanager.MonitoredTxStatusConfirmed {
				flag = false
			}
		})
	}
	etm.Stop()
	return nil
}
