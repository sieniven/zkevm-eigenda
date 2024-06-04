package main

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/sieniven/zkevm-eigenda/dataavailability/eigenda"
	"github.com/sieniven/zkevm-eigenda/etherman/smartcontracts/eigendaverifier"
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

	// Get auth
	auth, err := etherMan.GetAuthByAddress(adminAddr)
	if err != nil {
		panic(err)
	}

	// Estimate gas
	parsed, err := eigendaverifier.EigendaverifierMetaData.GetAbi()
	if err != nil || parsed == nil {
		return fmt.Errorf("estimate gas error, cannot get abi")
	}
	input, err := parsed.Pack("", adminAddr, c.L1Config.EigenDaServiceManagerAddr)
	if err != nil {
		panic(err)
	}
	bytecode := common.FromHex(eigendaverifier.EigendaverifierBin)
	input = append(bytecode, input...)
	gas, err := etherMan.EstimateGas(cliCtx.Context, adminAddr, nil, big.NewInt(0), input)
	if err != nil {
		panic(err)
	}
	fmt.Println("estimate gas: ", gas)

	// Connect to ethereum node
	ethClient, err := ethclient.Dial(c.Etherman.URL)
	if err != nil {
		fmt.Printf("error connecting to %s: %+v\n", c.Etherman.URL, err)
		return err
	}
	deployedAddr, deploymentTx, _, err := eigendaverifier.DeployEigendaverifier(&auth, ethClient, adminAddr, c.L1Config.EigenDaServiceManagerAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deployed to address: ", deployedAddr.String())
	fmt.Println("Deployment tx hash: ", deploymentTx.Hash().Hex())
	fmt.Println("Deployment tx info: ", deploymentTx)
	return nil
}
