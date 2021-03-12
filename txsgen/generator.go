package main

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TxMaker func(*ethclient.Client) (*types.Transaction, error)
type TxCallback func(*types.Receipt, error)

type Transaction struct {
	Make     TxMaker
	Callback TxCallback
	Dsc      string
}

type Generator interface {
	Start() (output <-chan *Transaction)
	Stop()
	SetTPS(tps float64)
}
