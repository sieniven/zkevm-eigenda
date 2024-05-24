package main

import (
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/urfave/cli/v2"
)

func testDAProvider(cliCtx *cli.Context) error {
	c, err := config.Load(cliCtx)
	if err != nil {
		return err
	}
	setupLog(c.Log)

	// da := dataavailability.New(c.EigenDAClient)

	// Generate mock string batch data
	data := []byte("hihihihihihihihihihihihihihihihihihi")

	// Generate mock string sequence
	return nil
}
