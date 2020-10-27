package presentation

import (
	"fmt"
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

	// DDL
	_, err = session.WriteTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		err := exec(ctx, "CREATE CONSTRAINT ON (e:Event) ASSERT e.Hash IS UNIQUE")
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		log.Warn(err.Error())
	}

	// DML
	var (
		counter int
		last    hash.Event
	)
	for event := range events {
		_, err = session.WriteTransaction(func(ctx neo4j.Transaction) (interface{}, error) {

			header := Neo4jMarshal(&event.EventHeaderData)
			log.Info("<<<", "event", header.String())
			err = exec(ctx, "CREATE (e:Event %s)", header.String())
			if err != nil {
				return nil, err
			}

			for _, p := range event.Parents {
				err = exec(ctx, `MATCH (e:Event {Hash:'%s'}), (p:Event {Hash:'%s'}) CREATE (e)-[:PARENT]->(p)`,
					event.Hash().Hex(),
					p.Hex(),
				)
				if err != nil {
					return nil, err
				}
			}

			return nil, nil

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

type executor interface {
	Run(cypher string, params map[string]interface{}) (neo4j.Result, error)
}

func exec(ctx executor, cypher string, a ...interface{}) error {
	query := fmt.Sprintf(cypher, a...)
	log.Debug("cypher", "query", query)
	res, err := ctx.Run(query, nil)
	if err != nil {
		return err
	}

	return res.Err()
}
