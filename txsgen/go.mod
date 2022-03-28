module github.com/Fantom-foundation/lachesis-dag-tool/txsgen

go 1.15

require (
	github.com/Fantom-foundation/go-opera v0.0.0-20220311122812-8d3ee1a19205
	github.com/Fantom-foundation/lachesis-base v0.0.0-20220103160934-6b4931c60582
	github.com/ethereum/go-ethereum v1.10.8
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	gopkg.in/urfave/cli.v1 v1.22.1
)

replace gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.20.0

replace github.com/ethereum/go-ethereum => github.com/Fantom-foundation/go-ethereum v1.10.8-ftm-rc4
