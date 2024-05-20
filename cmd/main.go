package main

import (
	"github.com/sieniven/polygoncdk-eigenda/config"
	"github.com/urfave/cli/v2"
)

const appNmae = "mock-zkevm-node"

const (
	// SEQUENCER is the sequencer component identifier
	SEQUENCER = "sequencer"
	// ETHTXMANAGER is the service that managers the tx sent to the L1
	ETHTXMANAGER = "eth-tx-manager"
	// SEQUENCE_SENDER is the sequence sender component identifier
	SEQUENCE_SENDER = "sequence-sender"
)

const (
	// NODE_CONFIGFILE name to identify the node config-file
	NODE_CONFIGFILE = "node"
)

var configFileFlag = cli.StringFlag{
	Name:     config.FlagCfg,
	Aliases:  []string{"c"},
	Usage:    "Configuration `FILE`",
	Required: true,
}

func main() {
	app := cli.NewApp()
	app.Name = appNmae
	flags := []cli.Flag{&configFileFlag}
	app.Commands = []*cli.Command{
		{
			Name:    "run",
			Aliases: []string{},
			Usage:   "Run the mock zkevm-node",
			Action:  start,
			Flags:   flags,
		},
		{
			Name:    "test-da",
			Aliases: []string{},
			Usage:   "Test the EigenDA client functionality",
			Action:  testEigenDA,
			Flags:   flags,
		},
	}
}
