package main

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/sieniven/zkevm-eigenda/etherman"
	"github.com/sieniven/zkevm-eigenda/ethtxmanager"
	"github.com/urfave/cli/v2"
)

func testEthTxManager(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return err
	}
	setupLog(c.Log)

	// Initialize ether manager instance
	etherMan, err := etherman.NewClient(c.Etherman, c.L1Config)
	if err != nil {
		panic(err)
	}

	// Create new data avaiability manager
	da := dataavailability.New(c.EigenDAClient)
	etherMan.SetDataProvider(da)

	// Initialize keys
	_, privKey, err := etherMan.LoadAuthFromKeyStore(c.Key.Path, c.Key.Password)
	if err != nil {
		panic(err)
	}
	if privKey == nil { //nolint:staticcheck
		panic(errors.New("private key not found"))
	}
	fmt.Printf("from pk %s, from sender %s\n", crypto.PubkeyToAddress(privKey.PublicKey), c.SequenceSender.SenderAddress.String()) //nolint:staticcheck

	// Get from address
	from := c.SequenceSender.SenderAddress

	// Log current balance of from address
	balance, err := etherMan.BalanceAt(cliCtx.Context, from)
	if err != nil {
		panic(err)
	}
	fmt.Println("Current account balance is: ", balance.Int64())

	// Start eth-tx-manager service
	etm := ethtxmanager.New(c.EthTxManager, etherMan)
	go etm.Start()

	// Get to address
	toStr := cliCtx.String(config.FlagTo)
	to := common.HexToAddress(toStr)

	// Test pipeline to create a monitored tx
	id := "test"
	owner := "sequencer"
	value := big.NewInt(10000000)
	etm.Add(cliCtx.Context, owner, id, from, &to, value, []byte{}, c.SequenceSender.GasOffset)

	// Test pipeline to test that monitored tx is sent successfully
	flag := true
	for flag {
		etm.ProcessPendingMonitoredTxs(cliCtx.Context, owner, func(result ethtxmanager.MonitoredTxResult) {
			if result.Status == ethtxmanager.MonitoredTxStatusFailed {
				fmt.Println("failed to send tx")
				flag = false
			}
		})
	}
	return nil
}
