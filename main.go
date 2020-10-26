package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/ethereum/go-ethereum/params"
	"github.com/urfave/cli"

	"github.com/Fantom-foundation/lachesis-dag-tool/presentation"
	"github.com/Fantom-foundation/lachesis-dag-tool/source"
)

var (
	// Git SHA1 commit hash of the release (set via linker flags).
	gitCommit = ""
	gitDate   = ""
	// App that holds all commands and flags.
	App = cli.NewApp()

	dataDirFlag = cli.StringFlag{
		Name:  "datadir",
		Usage: "Data directory of lachesis node",
		Value: source.DefaultDataDir(),
	}
	neo4jUrlFlag = cli.StringFlag{
		Name:  "neo4j",
		Usage: "Neo4j DB url",
		Value: presentation.Neo4jDefaultDb,
	}
)

// version of the current release
func version() string {
	params.VersionMajor = 0
	params.VersionMinor = 1
	params.VersionPatch = 0
	params.VersionMeta = "rc.1"
	return params.VersionWithCommit(gitCommit, gitDate)
}

func init() {
	App.Version = version()
	App.Flags = []cli.Flag{
		dataDirFlag,
		neo4jUrlFlag,
	}
	App.Action = cmdDefault
}

func cmdDefault(cli *cli.Context) (err error) {
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

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		events := source.Events(ctx, src, from, to)
		err = presentation.LoadToNeo4j(dst, events)
	}()

	waitForInterrupt(ctx)
	return
}

func waitForInterrupt(ctx context.Context) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sigs:
	case <-ctx.Done():
	}
}

func main() {
	if err := App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
