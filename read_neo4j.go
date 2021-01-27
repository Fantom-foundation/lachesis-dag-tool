package main

import (
	"context"
	"fmt"
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
		Action: cmd(actReadNeo4j),
		Usage:  "Read DAG from Neo4j db to compare performance with KV db.",
	}
)

func actReadNeo4j(ctx context.Context, cli *cli.Context) error {
	src := cli.String(neo4jUrlFlag.Name)
	if src == "" {
		src = cli.GlobalString(neo4jUrlFlag.Name)
	}

	event := hash.HexToEventHash(cli.Args().First())
	if event.IsZero() {
		err := fmt.Errorf("arg0 (event hash) required")
		return err
	}

	store, err := neo4j.New(src)
	if err != nil {
		return err
	}
	defer store.Close()

	start := time.Now()
	log.Info("Data from Neo4j", "database", src, "event", event)
	var ancestors []hash.Event
	ancestors = store.FindAncestors(event)
	log.Info("Data from Neo4j", "ancestors", len(ancestors), "elapsed", common.PrettyDuration(time.Since(start)))
	return nil
}
