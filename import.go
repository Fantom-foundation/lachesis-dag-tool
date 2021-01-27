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
		Action: cmd(actImport),
		Usage:  "Import DAG from lachesis datadir into Neo4j.",
	}

	networkFlag = cli.StringFlag{
		Name:  "network",
		Usage: "Lachesis network name (main or test)",
		Value: "main",
	}

	cmdListen = cli.Command{
		Name:      "listen",
		ShortName: "l",
		Flags: []cli.Flag{
			networkFlag,
		},
		Action: cmd(actListen),
		Usage:  "Listen lachesis p2p network and write DAG into Neo4j.",
	}
)

func actImport(ctx context.Context, cli *cli.Context) (err error) {
	dst := cli.GlobalString(neo4jUrlFlag.Name)
	src := cli.String(dataDirFlag.Name)
	from, to, err := parseEpochArgs(cli)
	if err != nil {
		return
	}

	store, err := neo4j.New(dst)
	if err != nil {
		return
	}
	defer store.Close()

	events := source.EventsFromDatadir(ctx, src, from, to, store)
	store.Load(events)
	return nil
}

func actListen(ctx context.Context, cli *cli.Context) error {
	dst := cli.GlobalString(neo4jUrlFlag.Name)
	network := cli.String(networkFlag.Name)
	from, to, err := parseEpochArgs(cli)
	if err != nil {
		return err
	}

	store, err := neo4j.New(dst)
	if err != nil {
		return err
	}
	defer store.Close()

	events := source.EventsFromP2p(ctx, network, from, to, store)
	store.Load(events)
	return nil
}

func parseEpochArgs(cli *cli.Context) (from, to idx.Epoch, err error) {
	var n uint64

	from = idx.Epoch(1)
	if len(cli.Args()) > 0 {
		n, err = strconv.ParseUint(cli.Args().Get(0), 10, 32)
		if err != nil {
			return
		}
		from = idx.Epoch(n)
	}

	to = idx.Epoch(0)
	if len(cli.Args()) > 1 {
		n, err = strconv.ParseUint(cli.Args().Get(1), 10, 32)
		if err != nil {
			return
		}
		to = idx.Epoch(n)
	}

	return
}
