package main

import (
	"context"
	"strconv"

	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/urfave/cli"

	"github.com/Fantom-foundation/lachesis-dag-tool/neo4j"
	"github.com/Fantom-foundation/lachesis-dag-tool/source"
)

var (
	dataDirFlag = cli.StringFlag{
		Name:  "datadir",
		Usage: "Data directory of lachesis node",
		Value: source.DefaultDataDir(),
	}

	cmdImport = cli.Command{
		Name:      "import",
		ShortName: "i",
		Flags: []cli.Flag{
			dataDirFlag,
		},
		Action: actImport,
		Usage:  "Import DAG from lachesis datadir into Neo4j.",
	}
)

func actImport(cli *cli.Context) (err error) {
	src := cli.GlobalString(dataDirFlag.Name)
	dst := cli.GlobalString(neo4jUrlFlag.Name)
	from := idx.Epoch(1)
	if len(cli.Args()) > 1 {
		n, err := strconv.ParseUint(cli.Args().Get(1), 10, 32)
		if err != nil {
			return err
		}
		from = idx.Epoch(n)
	}
	to := idx.Epoch(0)
	if len(cli.Args()) > 2 {
		n, err := strconv.ParseUint(cli.Args().Get(2), 10, 32)
		if err != nil {
			return err
		}
		to = idx.Epoch(n)
	}

	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer stop()
		events := source.Events(ctx, src, from, to)
		err = neo4j.LoadTo(dst, events)
	}()

	waitForInterrupt(ctx)
	return
}
