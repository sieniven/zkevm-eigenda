package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/sieniven/zkevm-eigenda/dataavailability"
	"github.com/sieniven/zkevm-eigenda/etherman"
	"github.com/sieniven/zkevm-eigenda/ethtxmanager"
	"github.com/sieniven/zkevm-eigenda/sequencesender"
	"github.com/urfave/cli/v2"
)

func start(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return err
	}

	// Initialize eth manager instance
	etherMan, err := newEtherman(*c)
	if err != nil {
		panic(err)
	}

	// Initialize eth tx manager instance
	etm := ethtxmanager.New(c.EthTxManager, etherMan)

	// Initialize mock sequence sender
	seqSender := createMockSequenceSender(*c, etm, etherMan)

	// Start mock sequence sender
	go seqSender.Start(cliCtx.Context)

	// Start send sequence flag handler
	reader := bufio.NewReader(os.Stdin)
	seqSender.SendSequenceHandle(cliCtx.Context, reader)
	return nil
}

func newEtherman(c config.Config) (*etherman.Client, error) {
	return etherman.NewClient(c.Etherman, c.L1Config)
}

// createMockSequenceSender is the mock function for PolygonCDK node that
// creates a new instance of the mock sequence sender for the mock node.
func createMockSequenceSender(cfg config.Config, etm *ethtxmanager.Client, etherMan *etherman.Client) *sequencesender.SequenceSender {
	da := setEthermanDA(cfg, etherMan)
	_, privKey, err := etherMan.LoadAuthFromKeyStore(cfg.Key.Path, cfg.Key.Password)
	if err != nil {
		panic(err)
	}
	if cfg.SequenceSender.SenderAddress.Cmp(common.Address{}) == 0 {
		panic(errors.New("sequence sender address not found"))
	}
	if privKey == nil { //nolint:staticcheck
		panic(errors.New("private key not found"))
	}
	fmt.Printf("from pk %s, from sender %s\n", crypto.PubkeyToAddress(privKey.PublicKey), cfg.SequenceSender.SenderAddress.String()) //nolint:staticcheck

	// Initialize new sequence sender instance
	seqSender, err := sequencesender.New(cfg.SequenceSender, etherMan, etm)
	if err != nil {
		panic(err)
	}
	seqSender.SetDataProvider(da)
	return seqSender
}

func newDataAvailability(c config.Config) (*dataavailability.DataAvailabilityProvider, error) {
	// TODO: add DA initialization pipline
	return &dataavailability.DataAvailabilityProvider{}, nil
}

func setEthermanDA(c config.Config, etherMan *etherman.Client) *dataavailability.DataAvailabilityProvider {
	da, err := newDataAvailability(c)
	if err != nil {
		panic(err)
	}
	etherMan.SetDataProvider(da)
	return da
}
