package main

import (
	"context"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/Fantom-foundation/lachesis-dag-tool/dagreader/neo4j"
)

var (
	neo4jUrlFlag = cli.StringFlag{
		Name:  "neo4j",
		Usage: "Neo4j DB url",
		Value: neo4j.DefaultDb,
	}

	cmdSaveTo = cli.Command{
		Name: "saveto",
		Flags: []cli.Flag{
			neo4jUrlFlag,
		},
		Action: cmd(actSaveTo),
		Usage:  "Write DAG into db.",
	}
)

func actSaveTo(ctx context.Context, cli *cli.Context) error {
	disk := cli.String(neo4jUrlFlag.Name)
	log.Info("open DB", "path", disk)
	db, err := neo4j.New(disk)
	if err != nil {
		return err
	}
	defer db.Close()

	buffer := NewEventsBuffer(db, ctx.Done())
	defer buffer.Close()

	rpc := cli.GlobalString(operaApiUrlFlag.Name)
	dagStart := idx.Block(cli.GlobalUint64(dagStartFlag.Name))

	log.Info("connect to API", "url", rpc)
	reader := NewReader(rpc, dagStart, db)
	defer reader.Close()

	for {
		select {
		case e := <-reader.Events():
			buffer.Push(e)
		case <-ctx.Done():
			return nil
		}
	}

	return nil
}
