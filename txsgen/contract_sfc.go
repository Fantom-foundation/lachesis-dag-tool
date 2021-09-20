package main

//go:generate bash -c "docker run --rm -v $(pwd)/ballot:/src -v $(pwd)/ballot:/dst ethereum/solc:0.5.12 -o /dst/ --optimize --optimize-runs=2000 --bin --abi --overwrite /src/Ballot.sol"
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --bin=./ballot/Ballot.bin --abi=./ballot/Ballot.abi --pkg=ballot --type=Contract --out=ballot/contract.go

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Fantom-foundation/lachesis-dag-tool/txsgen/ballot"
)

func (g *ReadonlyGenerator) ballotCreateContract(admin accounts.Account) TxMaker {
	payer := g.Payer(admin)
	return func(client *ethclient.Client) (*types.Transaction, error) {
		_, tx, _, err := ballot.DeployContract(payer, client, ballotOptions)
		if err != nil {
			panic(err)
		}

		return tx, err
	}
}

func (g *ReadonlyGenerator) ballotCountOfVoites(voiter accounts.Account, contract common.Address) TxMaker {
	payer := g.Payer(voiter, big.NewInt(100))
	return func(client *ethclient.Client) (*types.Transaction, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		opts := &bind.FilterOpts{
			Context: ctx,
		}
		filterer, err := ballot.NewContractFilterer(contract, client)
		if err != nil {
			panic(err)
		}
		logs, err := filterer.FilterVoiting(opts, []common.Address{contract, payer.From}, nil, nil)
		if err != nil {
			g.Log.Error("filterer.FilterVoiting()", "err", err)
			return nil, nil
		}
		defer logs.Close()

		var count int
		for ; logs.Next(); count++ {
		}
		g.Log.Info("prev voites", "count", count)

		return nil, nil
	}
}

func (g *ReadonlyGenerator) ballotVoite(voiter accounts.Account, contract common.Address, proposal int64) TxMaker {
	payer := g.Payer(voiter, big.NewInt(100))
	return func(client *ethclient.Client) (*types.Transaction, error) {
		transactor, err := ballot.NewContractTransactor(contract, client)
		if err != nil {
			panic(err)
		}

		return transactor.Vote(payer, big.NewInt(proposal))
	}
}

func (g *ReadonlyGenerator) ballotWinner(contract common.Address) TxMaker {
	return func(client *ethclient.Client) (*types.Transaction, error) {
		caller, err := ballot.NewContractCaller(contract, client)
		if err != nil {
			panic(err)
		}

		winner, err := caller.WinnerName(g.ReadOnly())
		g.Log.Info("The winner", "hash", winner)

		return nil, err
	}
}
