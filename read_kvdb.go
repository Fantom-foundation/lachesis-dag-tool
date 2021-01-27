package main

import (
	"context"
	"fmt"
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
		Action: cmd(actReadKVdb),
		Usage:  "Read DAG from KV db to compare performance with Neo4j db.",
	}
)

func actReadKVdb(ctx context.Context, cli *cli.Context) (err error) {
	src := cli.String(dataDirFlag.Name)

	event := hash.HexToEventHash(cli.Args().First())
	if event.IsZero() {
		err = fmt.Errorf("arg0 (event hash) required")
		return
	}

	start := time.Now()
	log.Info("Data from KV", "database", src, "event", event)
	var ancestors []hash.Event
	ancestors, err = source.FindAncestors(src, event)
	if err != nil {
		return
	}
	log.Info("Data from KV", "ancestors", len(ancestors), "elapsed", common.PrettyDuration(time.Since(start)))
	return
}
