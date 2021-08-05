package main

import (
	"errors"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/urfave/cli.v1"

	"github.com/Fantom-foundation/lachesis-dag-tool/txsgen/utils/toml"
)

var ConfigFileFlag = cli.StringFlag{
	Name:  "config",
	Usage: "TOML configuration file",
	Value: "txsgen.toml",
}

type Config struct {
	ChainId int64 // chain id for sign transactions
	Payer   common.Address
	URLs    []string // WS nodes API URL
}

func DefaultConfig() *Config {
	return &Config{
		ChainId: int64(opera.FakeNetworkID),
		URLs: []string{
			"ws://127.0.0.1:4500",
		},
	}
}

func OpenConfig(ctx *cli.Context) *Config {
	cfg := DefaultConfig()
	f := ctx.GlobalString(ConfigFileFlag.Name)
	err := cfg.Load(f)
	if err != nil {
		panic(err)
	}
	return cfg
}

func (cfg *Config) Load(file string) error {
	data, err := toml.ParseFile(file)
	if err != nil {
		return err
	}

	err = toml.Settings.UnmarshalTable(data, cfg)
	if err != nil {
		err = errors.New(file + ", " + err.Error())
		return err
	}

	return nil
}
