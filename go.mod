module github.com/Fantom-foundation/lachesis-dag-tool

go 1.14

require (
	github.com/Fantom-foundation/go-lachesis v0.8.0-rc.2
	github.com/ethereum/go-ethereum v1.9.8
	github.com/neo4j/neo4j-go-driver v1.8.3
	github.com/stretchr/testify v1.6.1
	github.com/urfave/cli v1.22.1
)

replace (
	github.com/ethereum/go-ethereum => github.com/Fantom-Foundation/go-ethereum v1.9.22-ftm-0.1
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.20.0
)
