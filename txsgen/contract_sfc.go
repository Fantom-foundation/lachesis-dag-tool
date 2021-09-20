package main

// compile SFC with truffle
//go:generate bash -c "cd ../../opera-sfc && git checkout c1d33c81f74abf82c0e22807f16e609578e10ad8"
//go:generate bash -c "docker run --name go-opera-sfc-compiler -v $(pwd)/sfc:/src/build/contracts -v $(pwd)/../../opera-sfc:/src -w /src node:10.5.0 bash -c 'export NPM_CONFIG_PREFIX=~; npm install --no-save; npm install --no-save truffle@5.1.4' && docker commit go-opera-sfc-compiler go-opera-sfc-compiler-image && docker rm go-opera-sfc-compiler"
//go:generate bash -c "docker run --rm -v $(pwd)/sfc:/src/build/contracts -v $(pwd)/../../opera-sfc:/src -w /src go-opera-sfc-compiler-image bash -c 'export NPM_CONFIG_PREFIX=~; rm -f /src/build/contracts/*json; npm run build'"
//go:generate bash -c "cd ./sfc && for f in LegacySfcWrapper.json; do jq -c .abi $DOLLAR{f} > $DOLLAR{f%.json}.abi; done"
//go:generate bash -c "docker rmi go-opera-sfc-compiler-image"
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi=./sfc/LegacySfcWrapper.abi --pkg=sfc --type=Contract --out=sfc/contract.go

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Fantom-foundation/lachesis-dag-tool/txsgen/sfc"
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
