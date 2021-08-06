package main

import (
	"context"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/Fantom-foundation/lachesis-dag-tool/dagreader/neo4j"
)

var (
	rpcFlag = cli.StringFlag{
		Name:  "rpc",
		Usage: "go-opera RPC url",
		Value: "ws://127.0.0.1:4500",
	}

	cmdListen = cli.Command{
		Name:      "listen",
		ShortName: "l",
		Flags: []cli.Flag{
			rpcFlag,
		},
		Action: cmd(actListen),
		Usage:  "Listen go-opera events and write DAG into Neo4j.",
	}
)

func actListen(ctx context.Context, cli *cli.Context) error {
	dst := cli.GlobalString(neo4jUrlFlag.Name)
	log.Info("open DB", "path", dst)
	disk, err := neo4j.New(dst)
	if err != nil {
		return err
	}
	db := neo4j.NewCachedDb(disk)
	defer db.Close()

	store := newStore(db, true)
	defer store.Close()

	// TODO: read from db
	var fromBlock idx.Block = 2

	rpc := cli.String(rpcFlag.Name)
	log.Info("connect to API", "url", rpc)
	reader := NewReader(rpc, fromBlock)
	defer reader.Close()

	for {
		select {
		case e := <-reader.Events():
			store.Save(e)
		case <-ctx.Done():
			return nil
		}
	}

	return nil
}
