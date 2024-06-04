package main

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/sieniven/zkevm-eigenda/dataavailability/eigenda"
	"github.com/sieniven/zkevm-eigenda/etherman/smartcontracts/eigendaverifier"
	"github.com/sieniven/zkevm-eigenda/ethtxmanager"
	"github.com/urfave/cli/v2"
)

func deployVerifier(cliCtx *cli.Context) error {
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

	admin := cliCtx.String(config.FlagAdmin)
	adminAddr := common.HexToAddress(admin)
	fmt.Println("Deploying with admin address: ", adminAddr)
	fmt.Printf("from pk %s, from sender %s\n", crypto.PubkeyToAddress(privKey.PublicKey), adminAddr.String()) //nolint:staticcheck
	fmt.Println("EigenDA service manager address: ", c.L1Config.EigenDaServiceManagerAddr)

	parsed, err := eigendaverifier.EigendaverifierMetaData.GetAbi()
	if err != nil || parsed == nil {
		return fmt.Errorf("estimate gas error, cannot get abi")
	}
	input, err := parsed.Pack("", adminAddr, c.L1Config.EigenDaServiceManagerAddr)
	if err != nil {
		panic(err)
	}

	// Get deployed EigenDARollupUtils lib address
	utilsAddr := c.L1Config.EigenDARollupUtilsAddr.Hex()[2:]
	paddedUtilsAddr := padWithZeros(utilsAddr, 40)

	// Replace placeholder with padded utils address
	placeholder := "__$399f9ce8dd33a06d144ce1eb24d845e280$__"
	bytecode := common.FromHex(strings.Replace(eigendaverifier.EigendaverifierBin, placeholder, paddedUtilsAddr, -1))

	input = append(bytecode, input...)
	gas, err := etherMan.EstimateGas(cliCtx.Context, adminAddr, nil, big.NewInt(0), input)
	if err != nil {
		panic(err)
	}
	fmt.Println("estimate gas: ", gas)

	// Start eth-tx-manager service
	etm := ethtxmanager.New(c.EthTxManager, etherMan)
	go etm.Start()

	// Send library deployment tx
	id := "deploy"
	owner := "deployment"
	value := big.NewInt(0)
	err = etm.Add(cliCtx.Context, owner, id, adminAddr, nil, value, input, c.SequenceSender.GasOffset)
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

func padWithZeros(input string, length int) string {
	if len(input) >= length {
		return input
	}
	padding := strings.Repeat("0", length-len(input))
	return input + padding
}
