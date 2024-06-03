package main

import (
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/urfave/cli/v2"
)

func deployVerifier(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return nil
	}
	setupLog((c.Log))

	// Connect to ethereum node
}
