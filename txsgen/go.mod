module github.com/Fantom-foundation/lachesis-dag-tool/txsgen

go 1.15

require (
	github.com/Fantom-foundation/go-lachesis v0.7.0-rc.1
	github.com/elastic/gosigar v0.10.5 // indirect
	github.com/ethereum/go-ethereum v1.9.8
	github.com/huin/goupnp v1.0.0 // indirect
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20200619000410-60c24ae608a6 // indirect
	gopkg.in/urfave/cli.v1 v1.22.1
)

replace gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.20.0

replace github.com/ethereum/go-ethereum => github.com/Fantom-foundation/go-ethereum v1.9.8-ftm-0.10
