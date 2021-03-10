package main

import (
	"gopkg.in/urfave/cli.v1"
)

var TxnsRateFlag = cli.IntFlag{
	Name:  "rate",
	Usage: "transactions per second (max sum of all instances)",
}

func getTxnsRate(ctx *cli.Context) uint {
	return uint(ctx.GlobalInt(TxnsRateFlag.Name))
}

var KeyStoreDirFlag = cli.StringFlag{
	Name:  "keystore",
	Usage: "Directory for the keystore",
	Value: "keystore",
}

var VerbosityFlag = cli.IntFlag{
	Name:  "verbosity",
	Usage: "sets the verbosity level",
	Value: 3,
}
