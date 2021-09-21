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
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Fantom-foundation/lachesis-dag-tool/txsgen/sfc"
)

var sfcContractAddress = common.HexToAddress("0xfc00face00000000000000000000000000000000")

func (g *ReadonlyGenerator) randomSfcCall() TxMaker {
	return func(client *ethclient.Client) (_ *types.Transaction, err error) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		block, err := client.BlockNumber(ctx)
		if err != nil {
			return
		}

		x := rand.Uint64() % 2
		if block > x {
			block -= x
		}
		ro := &bind.CallOpts{
			Context:     ctx,
			BlockNumber: big.NewInt(int64(block)),
		}

		caller, err := sfc.NewContractCaller(sfcContractAddress, client)
		if err != nil {
			panic(err)
		}

		var val *big.Int
		switch x {
		case 0:
			val, err = caller.ContractCommission(ro)
			if err != nil {
				return
			}
			g.Log.Info("Read SFC: ", "comission", val)
		case 1:
			val, err = caller.DelegationsNum(ro)
			if err != nil {
				return
			}
			g.Log.Info("Read SFC", "delegations", val)
		default:
			panic(fmt.Errorf("undefined case %d", x))
		}

		return
	}
}
