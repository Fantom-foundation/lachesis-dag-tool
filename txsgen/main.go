package main

import (
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"syscall"

	"github.com/Fantom-foundation/go-lachesis/crypto"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"gopkg.in/urfave/cli.v1"

	_ "github.com/Fantom-foundation/go-lachesis/version"
)

var (
	// Git SHA1 commit hash of the release (set via linker flags).
	gitCommit = ""
	gitDate   = ""
	// The app that holds all commands and flags.
	app = utils.NewApp(gitCommit, gitDate, "the transactions generator CLI")

	flags []cli.Flag

	mainCfg *Config
)

// init the CLI app.
func init() {
	app.Version = params.VersionWithCommit(gitCommit, gitDate)

	app.Commands = []cli.Command{
		cli.Command{
			Action:      makeFakenetAccs,
			Name:        "fakeaccs",
			Usage:       "[offset=1 [count=1000]]",
			Description: `Generates <count> fakenet accounts starting from <offset>.`,
		},
		cli.Command{
			Action:      generateCalls,
			Name:        "calls",
			Usage:       "Generates a lot of smart contract and web3-API calls.",
			Description: `Note: uses fakenet accounts and deploys a fake contract.`,
		},
		cli.Command{
			Action:      generateTransfers,
			Name:        "transfers",
			Usage:       "Generates a lot of transfer transactions.",
			Description: `Note: uses fakenet accounts.`,
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Flags = append(app.Flags,
		KeyStoreDirFlag,
		ConfigFileFlag,
		TxnsRateFlag,
		utils.MetricsEnabledFlag,
		MetricsPrometheusEndpointFlag,
		VerbosityFlag,
	)

	app.Before = before
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func before(ctx *cli.Context) error {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(ctx.GlobalInt(VerbosityFlag.Name)))
	log.Root().SetHandler(glogger)

	SetupPrometheus(ctx)
	mainCfg = OpenConfig(ctx)

	return nil
}

// makeFakenetAccs action.
func makeFakenetAccs(ctx *cli.Context) error {
	var accsOffset int = 1
	if ctx.NArg() > 1 {
		i64, err := strconv.ParseUint(ctx.Args().Get(1), 10, 64)
		if err != nil {
			return err
		}
		accsOffset = int(i64)
	}
	var accsCount int = 1000
	if ctx.NArg() > 2 {
		i64, err := strconv.ParseUint(ctx.Args().Get(2), 10, 64)
		if err != nil {
			return err
		}
		accsCount = int(i64)
	}

	keyStore, err := makeKeyStore(ctx)
	if err != nil {
		return err
	}

	for i := accsOffset; i < (accsOffset + accsCount); i++ {
		key := crypto.FakeKey(i)
		// addr := crypto.PubkeyToAddress(key.PublicKey)
		_, err := keyStore.ImportECDSA(key, "")
		if err != nil {
			return err
		}
	}

	return nil
}

// generateCalls action.
func generateCalls(ctx *cli.Context) error {
	cfg := mainCfg

	keyStore, err := makeKeyStore(ctx)
	if err != nil {
		return err
	}

	generator := NewCallsGenerator(cfg, keyStore)
	defer generator.Stop()
	generator.SetName("CallsGen")

	err = generate(generator)
	return err
}

// generateTransfers action.
func generateTransfers(ctx *cli.Context) error {
	cfg := mainCfg

	keyStore, err := makeKeyStore(ctx)
	if err != nil {
		return err
	}

	generator := NewTransfersGenerator(cfg, keyStore)
	defer generator.Stop()
	generator.SetName("TransfersGen-%d")

	err = generate(generator)
	return err
}

// generate is the main generate loop.
func generate(generator Generator) error {
	cfg := mainCfg
	txs := generator.Start()

	nodes := NewNodes(cfg, txs)
	go func() {
		for tps := range nodes.TPS() {
			generator.SetTPS(tps + 50.0*float64(nodes.Count()))
		}
	}()

	waitForSignal()
	return nil
}

func waitForSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
}
