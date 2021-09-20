package main

import (
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"syscall"

	"github.com/Fantom-foundation/go-opera/flags"
	"github.com/Fantom-foundation/go-opera/integration/makegenesis"
	_ "github.com/Fantom-foundation/go-opera/version"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"gopkg.in/urfave/cli.v1"
)

var (
	// Git SHA1 commit hash of the release (set via linker flags).
	gitCommit = ""
	gitDate   = ""
	// The app that holds all commands and flags.
	app = flags.NewApp(gitCommit, gitDate, "the transactions generator CLI")

	mainCfg *Config
)

// init the CLI app.
func init() {
	app.Version = params.VersionWithCommit(gitCommit, gitDate)

	app.Commands = []cli.Command{
		cli.Command{
			Action:      importAcc,
			Name:        "importacc",
			Usage:       "<address> <encrypted keystore dir> [password]",
			Description: `Decripts and imports account by <address> from <encrypted keystore dir> into the keystore dir.`,
		},

		cli.Command{
			Action:      makeFakenetAccs,
			Name:        "fakeaccs",
			Usage:       "[offset=1 [count=1000]]",
			Description: `Generates <count> fakenet accounts starting from <offset> and saves them in the keystore dir.`,
		},
		cli.Command{
			Action:      initAccsBalances,
			Name:        "initbalance",
			Usage:       "[amount=1]",
			Description: `Pays <amount> from config.Payer to each other account in the keystore dir.`,
		},
		cli.Command{
			Action:      generateCalls,
			Name:        "calls",
			Description: `Deploys a fake Contract and generates a lot of calls behalf of accounts in the keystore dir (except config.Payer).`,
		},
		cli.Command{
			Action:      generateReadonly,
			Name:        "readonly",
			Description: `Generates a lot of readonly API eth_call to the SFC.`,
		},
		cli.Command{
			Action:      generateTransfers,
			Name:        "transfers",
			Description: `Generates a lot of transfer txs between accounts in the keystore dir (except config.Payer).`,
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Flags = append(app.Flags,
		KeyStoreDirFlag,
		ConfigFileFlag,
		TpsLimitFlag,
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

// importAcc action.
func importAcc(ctx *cli.Context) error {
	if ctx.NArg() < 2 {
		return fmt.Errorf("Address and Keystore dir args expected")
	}

	acc := accounts.Account{
		Address: common.HexToAddress(ctx.Args().Get(0)),
	}
	other, err := openKeyStore(ctx.Args().Get(1))
	if err != nil {
		return err
	}

	var password string
	if ctx.NArg() > 2 {
		password = ctx.Args().Get(2)
	}

	my, err := makeKeyStore(ctx)
	if err != nil {
		return err
	}

	decrypted, err := other.Export(acc, password, "")
	if err != nil {
		return err
	}

	_, err = my.Import(decrypted, "", "")
	if err != nil {
		return err
	}

	return nil
}

// makeFakenetAccs action.
func makeFakenetAccs(ctx *cli.Context) error {
	var accsOffset int = 1
	if ctx.NArg() > 0 {
		i64, err := strconv.ParseUint(ctx.Args().Get(0), 10, 64)
		if err != nil {
			return err
		}
		accsOffset = int(i64)
	}
	var accsCount int = 1000
	if ctx.NArg() > 1 {
		i64, err := strconv.ParseUint(ctx.Args().Get(1), 10, 64)
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
		key := makegenesis.FakeKey(i)
		_, err := keyStore.ImportECDSA(key, "")
		if err != nil {
			return err
		}
	}

	return nil
}

// initAccsBalances action.
func initAccsBalances(ctx *cli.Context) error {
	cfg := mainCfg
	cfg.URLs = cfg.URLs[:1] // txs from single payer should be sent by single sender
	keyStore, err := makeKeyStore(ctx)
	if err != nil {
		return err
	}

	var amount int64 = 1e18
	if ctx.NArg() > 0 {
		i64, err := strconv.ParseUint(ctx.Args().Get(0), 10, 64)
		if err != nil {
			return err
		}
		amount = int64(i64)
	}

	maxTps := getTpsLimit(ctx)

	generator := NewBalancesGenerator(cfg, keyStore, amount)
	generator.SetName("InitBalance")
	err = generate(generator, maxTps)
	return err
}

// generateCalls action.
func generateCalls(ctx *cli.Context) error {
	cfg := mainCfg
	keyStore, err := makeKeyStore(ctx)
	if err != nil {
		return err
	}

	maxTps := getTpsLimit(ctx)

	generator := NewCallsGenerator(cfg, keyStore)
	generator.SetName("CallsGen")
	err = generate(generator, maxTps)
	return err
}

// generateReadonly action.
func generateReadonly(ctx *cli.Context) error {
	cfg := mainCfg
	keyStore, err := makeKeyStore(ctx)
	if err != nil {
		return err
	}

	maxTps := getTpsLimit(ctx)

	generator := NewReadonlyGenerator(cfg, keyStore)
	generator.SetName("ReadonlyGen")
	err = generate(generator, maxTps)
	return err
}

// generateTransfers action.
func generateTransfers(ctx *cli.Context) error {
	cfg := mainCfg
	keyStore, err := makeKeyStore(ctx)
	if err != nil {
		return err
	}

	maxTps := getTpsLimit(ctx)

	generator := NewTransfersGenerator(cfg, keyStore)
	generator.SetName("TransfersGen")
	err = generate(generator, maxTps)
	return err
}

// generate is the main generate loop.
func generate(generator Generator, maxTps float64) error {
	cfg := mainCfg
	txs := generator.Start()
	defer generator.Stop()

	nodes := NewNodes(cfg, txs)
	go func() {
		for tps := range nodes.TPS() {
			tps += 10.0 * float64(nodes.Count())
			if maxTps > 0.0 && tps > maxTps {
				tps = maxTps
			}
			generator.SetTPS(tps)
		}
	}()

	waitForFinish(nodes.Done)
	return nil
}

func waitForFinish(done <-chan struct{}) {
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		break
	case <-done:
		break
	}
}
