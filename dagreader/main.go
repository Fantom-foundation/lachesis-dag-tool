package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
)

var (
	// App that holds all commands and flags.
	App = cli.NewApp()

	operaApiUrlFlag = cli.StringFlag{
		Name:  "api",
		Usage: "opera API url",
		Value: "ws://127.0.0.1:4500",
	}

	dagStartFlag = cli.Uint64Flag{
		Name:  "dagstart",
		Usage: "genesis blocks with no DAG to skip them (4564024 for mainnet)",
		Value: 1,
	}
)

func init() {
	log.Root().SetHandler(
		log.LvlFilterHandler(log.LvlInfo, log.StdoutHandler))

	App.Version = version()
	App.Flags = []cli.Flag{
		operaApiUrlFlag,
		dagStartFlag,
	}
	App.Commands = []cli.Command{
		cmdSaveTo,
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
