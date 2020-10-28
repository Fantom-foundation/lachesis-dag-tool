package main

import (
	"context"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
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
		log.Info("Data from KV", "database", src)
		// TODO: read data from KV database.
		err = nil
	}()

	waitForInterrupt(ctx)
	return
}
