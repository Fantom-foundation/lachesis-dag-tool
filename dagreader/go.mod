module github.com/Fantom-foundation/lachesis-dag-tool/dagreader

go 1.14

require (
	github.com/Fantom-foundation/go-opera v1.1.1-rc.2
	github.com/Fantom-foundation/lachesis-base v0.0.0-20220103160934-6b4931c60582
	github.com/ethereum/go-ethereum v1.10.8
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/neo4j/neo4j-go-driver v1.8.3
	github.com/paulbellamy/ratecounter v0.2.0
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.1-0.20210305035536-64b5b1c73954
	github.com/urfave/cli v1.22.1
	gopkg.in/urfave/cli.v1 v1.22.1 // indirect
)

replace (
	github.com/ethereum/go-ethereum => github.com/Fantom-foundation/go-ethereum v1.10.8-ftm-rc5
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.20.0
)
