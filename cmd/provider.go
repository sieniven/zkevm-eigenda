package main

import (
	"github.com/Layr-Labs/eigenda/encoding/utils/codec"
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
	data = codec.ConvertByPaddingEmptyByte(data)

	// Generate mock string sequence
	return nil
}
