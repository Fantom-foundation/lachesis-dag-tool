package main

import (
	"context"
	"time"

	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/Fantom-foundation/lachesis-dag-tool/source"
)

var (
	cmdReadKVdb = cli.Command{
		Name: "read1",
		Flags: []cli.Flag{
			dataDirFlag,
		},
		Action: actReadKVdb,
		Usage:  "Read DAG from KV db to compare performance with Neo4j db.",
	}
)

func actReadKVdb(cli *cli.Context) (err error) {
	src := cli.String(dataDirFlag.Name)

	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer stop()

		start := time.Now()
		log.Info("Data from KV", "database", src)
		// TODO: read event hash from cli arg.
		event := hash.HexToEventHash("0x0000000400000016f9a4c23827a98e8dfa1358a41eb79d71e889c97c973722ab")
		var ancestors []hash.Event
		ancestors, err = source.FindAncestors(src, event)
		if err != nil {
			return
		}
		log.Info("Data from KV", "ancestors", len(ancestors), "elapsed", common.PrettyDuration(time.Since(start)))
	}()

	waitForInterrupt(ctx)
	return
}
