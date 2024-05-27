package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/sieniven/zkevm-eigenda/config"
	"github.com/urfave/cli/v2"
)

func createKeystore(cliCtx *cli.Context) error {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := cliCtx.String(config.FlagPassword)

	account, err := ks.NewAccount(password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated account with address: ", account.Address.Hex())
	return nil
}
