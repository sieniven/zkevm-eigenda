package main

import (
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/urfave/cli/v2"
)

func testEthTxManager(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return err
	}
	setupLog(c.Log)
	return nil
}
