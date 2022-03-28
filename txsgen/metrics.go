package main

import (
	"net/http"

	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/metrics/prometheus"
	cli "gopkg.in/urfave/cli.v1"
)

var MetricsPrometheusEndpointFlag = cli.StringFlag{
	Name:  "metrics.prometheus.endpoint",
	Usage: "Prometheus API endpoint to report metrics to",
	Value: ":19090",
}

var (
	reg = metrics.NewRegistry()

	txCountSentMeter = metrics.NewRegisteredCounter("tx_count_sent", reg)
	txCountGotMeter  = metrics.NewRegisteredCounter("tx_count_got", reg)
	txTpsMeter       = metrics.NewRegisteredHistogram("tx_tps", reg, metrics.NewUniformSample(500))
)

func SetupPrometheus(ctx *cli.Context) {
	if !metrics.Enabled {
		return
	}

	var endpoint = ctx.GlobalString(MetricsPrometheusEndpointFlag.Name)

	go http.ListenAndServe(endpoint, prometheus.Handler(reg))
}
