package main

import (
	"context"
	"time"

	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/Fantom-foundation/lachesis-dag-tool/neo4j"
)

var (
	cmdReadNeo4j = cli.Command{
		Name: "read",
		Flags: []cli.Flag{
			neo4jUrlFlag,
		},
		Action: actReadNeo4j,
		Usage:  "Read DAG from Neo4j db to compare performance with KV db.",
	}
)

func actReadNeo4j(cli *cli.Context) (err error) {
	src := cli.String(neo4jUrlFlag.Name)
	if src == "" {
		src = cli.GlobalString(neo4jUrlFlag.Name)
	}

	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer stop()
		start := time.Now()
		log.Info("Data from Neo4j", "database", src)
		// TODO: read event hash from cli arg.
		event := hash.HexToEventHash("0x0000000400000016f9a4c23827a98e8dfa1358a41eb79d71e889c97c973722ab")
		var ancestors []hash.Event
		ancestors, err = neo4j.FindAncestors(src, event)
		if err != nil {
			return
		}
		log.Info("Data from Neo4j", "ancestors", len(ancestors), "elapsed", common.PrettyDuration(time.Since(start)))
	}()

	waitForInterrupt(ctx)
	return
}
