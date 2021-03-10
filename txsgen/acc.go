package main

import (
	"crypto/ecdsa"
	"math/big"
	"os"
	"path/filepath"

	"github.com/Fantom-foundation/go-lachesis/crypto"
	"github.com/Fantom-foundation/go-lachesis/lachesis/params"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/urfave/cli.v1"
)

var (
	gasLimit = uint64(21000)
	gasPrice = params.MinGasPrice // minimal
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

	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	keyStore := keystore.NewKeyStore(keydir, scryptN, scryptP)

	return keyStore, nil
}

type Acc struct {
	Key  *ecdsa.PrivateKey
	Addr *common.Address
}

func MakeAcc(n uint) *Acc {
	key := crypto.FakeKey(int(n))
	addr := crypto.PubkeyToAddress(key.PublicKey)

	return &Acc{
		Key:  key,
		Addr: &addr,
	}
}

func (a *Acc) TransactionTo(b *Acc, nonce uint, amount *big.Int, chainId uint) *types.Transaction {
	tx := types.NewTransaction(
		uint64(nonce),
		*b.Addr,
		amount,
		gasLimit,
		gasPrice,
		[]byte{},
	)

	signed, err := types.SignTx(
		tx,
		types.NewEIP155Signer(big.NewInt(int64(chainId))),
		a.Key,
	)
	if err != nil {
		panic(err)
	}

	return signed
}
