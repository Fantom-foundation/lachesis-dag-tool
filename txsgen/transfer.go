package main

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (g *TransfersGenerator) transferTx(from, to accounts.Account, nonce uint) TxMaker {
	amount := big.NewInt(1e6)

	tx := types.NewTransaction(
		uint64(nonce),
		to.Address,
		amount,
		gasLimit,
		gasPrice,
		[]byte{},
	)

	signed, err := g.accs.SignTx(from, tx, g.chainId)
	if err != nil {
		panic(err)
	}

	return func(client *ethclient.Client) (*types.Transaction, error) {
		err := client.SendTransaction(context.Background(), signed)
		return signed, err
	}
}
