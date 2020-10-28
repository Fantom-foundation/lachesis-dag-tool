package main

import (
	"context"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
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
		log.Info("Data from Neo4j", "database", src)
		// TODO: read data from Neo4j database.
		err = nil
	}()

	waitForInterrupt(ctx)
	return
}
