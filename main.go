package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/urfave/cli"

	"github.com/Fantom-foundation/lachesis-dag-tool/neo4j"
)

var (
	// Git SHA1 commit hash of the release (set via linker flags).
	gitCommit = ""
	gitDate   = ""
	// App that holds all commands and flags.
	App = cli.NewApp()

	neo4jUrlFlag = cli.StringFlag{
		Name:  "neo4j",
		Usage: "Neo4j DB url",
		Value: neo4j.DefaultDb,
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
	log.Root().SetHandler(
		log.LvlFilterHandler(log.LvlInfo, log.StdoutHandler))

	App.Version = version()
	App.Flags = []cli.Flag{
		neo4jUrlFlag,
	}
	App.Commands = []cli.Command{
		cmdImport,
		cmdReadNeo4j,
		cmdReadKVdb,
	}
}

func main() {
	if err := App.Run(os.Args); err != nil {
		log.Crit("Fail", "err", err)
	}
}

func cmd(f func(context.Context, *cli.Context) error) func(*cli.Context) error {
	return func(cli *cli.Context) error {
		ctx, cancel := jobContext()
		defer cancel()
		return f(ctx, cli)
	}
}

func jobContext() (ctx context.Context, cancel func()) {
	ctx, cancel = context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		select {
		case <-sigs:
			log.Warn("Interrupted")
			cancel()
		case <-ctx.Done():
			log.Info("Finished")
		}
	}()

	return
}
