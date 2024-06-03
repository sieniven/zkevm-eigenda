package main

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/sieniven/zkevm-eigenda/dataavailability/eigenda"
	"github.com/sieniven/zkevm-eigenda/etherman"
	"github.com/urfave/cli/v2"
)

func testEtherman(cliCtx *cli.Context) error {
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
	p := eigenda.NewDataProvider(c.DataAvailability)
	da := dataavailability.New(c.DataAvailability, p)
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

	// Get to address
	toStr := cliCtx.String(config.FlagTo)
	to := common.HexToAddress(toStr)

	// Set default value
	value := big.NewInt(999700000000000000)

	// Get gas
	gas, err := etherMan.EstimateGas(cliCtx.Context, from, &to, value, []byte{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Estimate gas is: ", gas)

	// get gas price
	gasPrice, err := etherMan.SuggestedGasPrice(cliCtx.Context)
	if err != nil {
		panic(err)
	}
	fmt.Println("The suggested gasPrice before margin factor is: ", gasPrice)
	// adjust the gas price by the margin factor
	marginFactor := big.NewFloat(0).SetFloat64(c.EthTxManager.GasPriceMarginFactor)
	fGasPrice := big.NewFloat(0).SetInt(gasPrice)
	gasPrice, _ = big.NewFloat(0).Mul(fGasPrice, marginFactor).Int(big.NewInt(0))

	// if there is a max gas price limit configured and the current
	// adjusted gas price is over this limit, set the gas price as the limit
	if c.EthTxManager.MaxGasPriceLimit > 0 {
		maxGasPrice := big.NewInt(0).SetUint64(c.EthTxManager.MaxGasPriceLimit)
		if gasPrice.Cmp(maxGasPrice) == 1 {
			gasPrice.Set(maxGasPrice)
		}
	}
	fmt.Println("The suggested gasPrice after margin factor is: ", gasPrice)

	// Get nonce
	nonce, err := etherMan.CurrentNonce(cliCtx.Context, c.SequenceSender.SenderAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("Nonce: ", nonce)

	// Sign test tx
	tx := types.NewTx(&types.LegacyTx{
		To:       &to,
		Nonce:    nonce,
		Value:    value,
		Gas:      gas,
		GasPrice: gasPrice,
	})
	signedtx, err := etherMan.SignTx(cliCtx.Context, from, tx)
	if err != nil {
		panic(err)
	}

	// Send test tx
	err = etherMan.SendTx(cliCtx.Context, signedtx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sent signed tx to node")

	return nil
}
