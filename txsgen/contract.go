package main

//go:generate bash -c "docker run --rm -v $(pwd)/ballot:/src -v $(pwd)/ballot:/dst ethereum/solc:0.5.12 -o /dst/ --optimize --optimize-runs=2000 --bin --abi --overwrite /src/Ballot.sol"
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --bin=./ballot/Ballot.bin --abi=./ballot/Ballot.abi --pkg=ballot --type=Contract --out=ballot/contract.go

import (
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/Fantom-foundation/go-lachesis/hash"

	"github.com/Fantom-foundation/lachesis-dag-tool/txsgen/ballot"
)

var ballotOptions = [][32]byte{
	ballotProposal("option 1"),
	ballotProposal("option 2"),
	ballotProposal("option 3"),
	ballotProposal("option 4"),
	ballotProposal("option 5"),
	ballotProposal("option 6"),
	ballotProposal("option 7"),
	ballotProposal("option 8"),
	ballotProposal("option 9"),
}

func ballotRandChose() int64 {
	return rand.Int63n(int64(len(ballotOptions)))
}

func (g *CallsGenerator) ballotCreateContract(admin uint) TxMaker {
	payer := g.Payer(admin)
	return func(client *ethclient.Client) (*types.Transaction, error) {
		_, tx, _, err := ballot.DeployContract(payer, client, ballotOptions)
		if err != nil {
			panic(err)
		}

		return tx, err
	}
}

func (g *CallsGenerator) ballotVoite(voiter uint, addr common.Address, proposal int64) TxMaker {
	payer := g.Payer(voiter)
	return func(client *ethclient.Client) (*types.Transaction, error) {
		transactor, err := ballot.NewContractTransactor(addr, client)
		if err != nil {
			panic(err)
		}

		return transactor.Vote(payer, big.NewInt(proposal))
	}
}

func (g *CallsGenerator) ballotWinner(addr common.Address) TxMaker {
	return func(client *ethclient.Client) (*types.Transaction, error) {
		caller, err := ballot.NewContractCaller(addr, client)
		if err != nil {
			panic(err)
		}

		winner, err := caller.WinnerName(g.ReadOnly())
		g.Log.Info("The winner", "hash", winner)

		return nil, err
	}
}

func ballotProposal(s string) [32]byte {
	return hash.Of([]byte(s))
}
