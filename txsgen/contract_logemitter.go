package main

//go:generate bash -c "docker run --rm -v $(pwd)/logemitter:/src -v $(pwd)/logemitter:/dst ethereum/solc:0.5.12 -o /dst/ --optimize --optimize-runs=2000 --bin --abi --overwrite /src/LogEmitter.sol"
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --bin=./logemitter/LogEmitter.bin --abi=./logemitter/LogEmitter.abi --pkg=logemitter --type=Contract --out=logemitter/contract.go

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Fantom-foundation/lachesis-dag-tool/txsgen/logemitter"
)

func (g *CallsGenerator) logemitterCreateContract(admin accounts.Account) TxMaker {
	payer := g.Payer(admin)
	return func(client *ethclient.Client) (*types.Transaction, error) {
		_, tx, _, err := logemitter.DeployContract(payer, client)
		if err != nil {
			panic(err)
		}

		return tx, err
	}
}
