package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (g *TransfersGenerator) transferTx(from, to *Acc, nonce uint) TxMaker {
	amount := big.NewInt(1e6)

	return func(client *ethclient.Client) (*types.Transaction, error) {
		tx := from.TransactionTo(to, nonce, amount, g.chainId)
		return tx, nil
	}
}
