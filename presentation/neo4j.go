package presentation

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/neo4j/neo4j-go-driver/neo4j"

	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter"
)

const (
	Neo4jDefaultDb = "bolt://localhost:7687"

	// statsReportLimit is the time limit during import and export after which we
	// always print out progress. This avoids the user wondering what's going on.
	statsReportLimit = 8 * time.Second
)

// exportTo writer the active chain.
func LoadToNeo4j(dbUrl string, events <-chan *inter.Event) (err error) {
	start, reported := time.Now(), time.Time{}

	db, err := neo4j.NewDriver(dbUrl, neo4j.NoAuth(), func(c *neo4j.Config) {
		c.Encrypted = false
	})
	if err != nil {
		return err
	}
	defer db.Close()

	session, err := db.Session(neo4j.AccessModeWrite)
	if err != nil {
		return
	}
	defer session.Close()

	var (
		encoder = NewNeo4jEncoding(&inter.EventHeaderData{}, "GasPowerLeft", "Parents")
		counter int
		last    hash.Event
	)
	for event := range events {
		_, err = session.WriteTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
			header, err := encoder.Marshal(&event.EventHeaderData)
			if err != nil {
				return nil, err
			}

			result, err := ctx.Run(
				"CREATE (e:Event "+string(header)+")",
				nil)
			if err != nil {
				panic(string(header))
				return nil, err
			}

			return nil, result.Err()

		})
		if err != nil {
			log.Error("<<<", "err", err)
			return err
		}

		counter++
		last = event.Hash()
		if counter%100 == 1 && time.Since(reported) >= statsReportLimit {
			log.Info("<<<", "last", last.String(), "exported", counter, "elapsed", common.PrettyDuration(time.Since(start)))
			reported = time.Now()
		}
	}

	log.Info("Exported events", "last", last.String(), "exported", counter, "elapsed", common.PrettyDuration(time.Since(start)))
	return
}
