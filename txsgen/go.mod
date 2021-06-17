module github.com/Fantom-foundation/lachesis-dag-tool/txsgen

go 1.15

require (
	github.com/Fantom-foundation/go-opera v0.0.0-20210505071830-e544db76285c
	github.com/Fantom-foundation/lachesis-base v0.0.0-20210420092627-c16f01e35562
	github.com/ethereum/go-ethereum v1.10.3
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	gopkg.in/urfave/cli.v1 v1.22.1
)

replace gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.20.0

replace github.com/ethereum/go-ethereum => github.com/Fantom-foundation/go-ethereum v1.9.7-0.20210531094457-b859cd9c4511

replace github.com/Fantom-foundation/go-opera => github.com/rus-alex/go-opera v0.0.0-20210617072707-c1a3e5a41464
