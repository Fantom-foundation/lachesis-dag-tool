package main

//go:generate bash -c "docker run --rm -v $(pwd)/logemitter:/src -v $(pwd)/logemitter:/dst ethereum/solc:0.5.12 -o /dst/ --optimize --optimize-runs=2000 --bin --abi --overwrite /src/LogEmitter.sol"
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --bin=./logemitter/LogEmitter.bin --abi=./logemitter/LogEmitter.abi --pkg=logemitter --type=Contract --out=logemitter/contract.go

import (
	"context"
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Fantom-foundation/lachesis-dag-tool/txsgen/logemitter"
)

func (g *GetLogsGenerator) logemitterCreateContract(admin accounts.Account) TxMaker {
	payer := g.Payer(admin)
	return func(client *ethclient.Client) (*types.Transaction, error) {
		_, tx, _, err := logemitter.DeployContract(payer, client)
		if err != nil {
			panic(err)
		}

		return tx, err
	}
}

func (g *GetLogsGenerator) logemitterEmit(sender accounts.Account, contract common.Address, seed uint) TxMaker {
	payer := g.Payer(sender, big.NewInt(100))
	return func(client *ethclient.Client) (*types.Transaction, error) {
		transactor, err := logemitter.NewContractTransactor(contract, client)
		if err != nil {
			panic(err)
		}

		rnd := rand.New(rand.NewSource(int64(seed)))
		var data [8][32]byte
		for i := range data {
			rnd.Read(data[i][:])
		}

		return transactor.EmitLogs(payer, data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7])
	}
}

func (g *GetLogsGenerator) logemitterGet(sender accounts.Account, contract common.Address, seed uint) TxMaker {
	return func(client *ethclient.Client) (*types.Transaction, error) {
		rnd := rand.New(rand.NewSource(int64(seed)))
		var data [8][]byte
		for i := range data {
			data[i] = make([]byte, 32)
			rnd.Read(data[i])
		}

		ctx := context.TODO()
		_, err := client.FilterLogs(ctx, ethereum.FilterQuery{
			Addresses: []common.Address{sender.Address},
			Topics: [][]common.Hash{
				[]common.Hash{common.BytesToHash(data[0])},
				[]common.Hash{common.BytesToHash(data[1])},
				[]common.Hash{common.BytesToHash(data[2])},
				[]common.Hash{common.BytesToHash(data[3])},
				[]common.Hash{common.BytesToHash(data[4])},
				[]common.Hash{common.BytesToHash(data[5])},
				[]common.Hash{common.BytesToHash(data[6])},
				[]common.Hash{common.BytesToHash(data[7])},
			},
		})
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
