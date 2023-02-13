package main

import (
	"math/big"
	"os"
	"path/filepath"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"gopkg.in/urfave/cli.v1"
)

var (
	gasLimit = uint64(21000)
	gasPrice = new(big.Int).Mul(opera.FakeEconomyRules().MinGasPrice, big.NewInt(2))
)

func makeKeyStore(ctx *cli.Context) (*keystore.KeyStore, error) {
	keydir := ctx.GlobalString(KeyStoreDirFlag.Name)
	keydir, err := filepath.Abs(keydir)
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(keydir, 0700)
	if err != nil {
		return nil, err
	}

	keyStore := keystore.NewPlaintextKeyStore(keydir)

	return keyStore, nil
}

func openKeyStore(keydir string) (*keystore.KeyStore, error) {
	keydir, err := filepath.Abs(keydir)
	if err != nil {
		return nil, err
	}

	keyStore := keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)

	return keyStore, nil
}
