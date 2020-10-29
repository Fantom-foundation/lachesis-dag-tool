package neo4j

import (
	"fmt"
	"time"

	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/paulbellamy/ratecounter"
)

const (
	DefaultDb = "bolt://localhost:7687"

	// statsReportLimit is the time limit during import and export after which we
	// always print out progress. This avoids the user wondering what's going on.
	statsReportLimit = 8 * time.Second
)

// LoadTo Neo4j from events chain.
func LoadTo(dbUrl string, events <-chan *inter.Event) error {
	db, err := neo4j.NewDriver(dbUrl, neo4j.NoAuth(), func(c *neo4j.Config) {
		c.Encrypted = false
	})
	if err != nil {
		return err
	}
	defer db.Close()

	session, err := db.Session(neo4j.AccessModeWrite)
	if err != nil {
		return err
	}
	defer session.Close()

	// DDL
	_, err = session.WriteTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		err := exec(ctx, "CREATE CONSTRAINT ON (e:Event) ASSERT e.hash IS UNIQUE")
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
		start    = time.Now().Add(-10 * time.Millisecond)
		reported time.Time
		counter  = ratecounter.NewRateCounter(60 * time.Second).WithResolution(1)
		total    int64
		last     hash.Event
	)
	for event := range events {
		_, err = session.WriteTransaction(func(ctx neo4j.Transaction) (interface{}, error) {

			header := Marshal(&event.EventHeaderData)
			log.Debug("<<<", "event", header.String())
			err = exec(ctx, "CREATE (e:Event %s)", header.String())
			if err != nil {
				_ = ctx.Rollback()
				return nil, err
			}

			for _, p := range event.Parents {
				err = exec(ctx, `MATCH (e:Event {hash:'%s'}), (p:Event {hash:'%s'}) CREATE (e)-[:PARENT]->(p)`,
					event.Hash().Hex(),
					p.Hex(),
				)
				if err != nil {
					_ = ctx.Rollback()
					return nil, err
				}
			}

			return nil, ctx.Commit()

		})
		if err != nil {
			log.Error("<<<", "err", err)
			// return err
		}

		counter.Incr(1)
		total++
		last = event.Hash()
		if time.Since(reported) >= statsReportLimit {
			log.Info("<<<", "last", last.String(),
				"per second", counter.Rate()/60,
				"total", total,
				"elapsed", common.PrettyDuration(time.Since(start)))
			reported = time.Now()
		}
	}

	log.Info("Exported events", "last", last.String(),
		"per second", total*1000/time.Since(start).Milliseconds(),
		"total", total,
		"elapsed", common.PrettyDuration(time.Since(start)))
	return nil
}

// FindAncestors of event.
func FindAncestors(dbUrl string, event hash.Event) (ancestors []hash.Event, err error) {
	db, err := neo4j.NewDriver(dbUrl, neo4j.NoAuth(), func(c *neo4j.Config) {
		c.Encrypted = false
	})
	if err != nil {
		return
	}
	defer db.Close()

	session, err := db.Session(neo4j.AccessModeRead)
	if err != nil {
		return
	}
	defer session.Close()

	res, err := session.ReadTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		res, err := search(ctx, "MATCH (p:Event {hash:'%s'})-[:PARENT*]->(s:Event) RETURN DISTINCT s.hash", event.Hex())
		if err != nil {
			return nil, err
		}

		var ancestors []hash.Event
		for res.Next() {
			hex := res.Record().GetByIndex(0).(string)
			ancestors = append(ancestors, hash.HexToEventHash(hex))
		}
		return ancestors, nil
	})
	if err != nil {
		return
	}

	ancestors = res.([]hash.Event)
	return
}

func exec(ctx neo4j.Transaction, cypher string, a ...interface{}) error {
	query := fmt.Sprintf(cypher, a...)
	log.Debug("cypher", "query", query)
	res, err := ctx.Run(query, nil)
	if err != nil {
		return err
	}

	return res.Err()
}

func search(ctx neo4j.Transaction, cypher string, a ...interface{}) (neo4j.Result, error) {
	query := fmt.Sprintf(cypher, a...)
	log.Debug("cypher", "query", query)
	res, err := ctx.Run(query, nil)
	if err != nil {
		return nil, err
	}

	return res, res.Err()
}
