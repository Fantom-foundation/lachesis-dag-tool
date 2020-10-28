package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/urfave/cli"

	"github.com/Fantom-foundation/lachesis-dag-tool/presentation"
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
		neo4jUrlFlag,
	}
	App.Commands = []cli.Command{
		cmdImport,
	}
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
		log.Crit("Fail", "err", err)
	}
}
