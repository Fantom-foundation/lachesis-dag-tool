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

type Store struct {
	db neo4j.Driver
}

func New(dbUrl string) (*Store, error) {
	db, err := neo4j.NewDriver(dbUrl, neo4j.NoAuth(), func(c *neo4j.Config) {
		c.Encrypted = false
	})
	if err != nil {
		return nil, err
	}

	session, err := db.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	// DDL
	_, err = session.WriteTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		err := exec(ctx, "CREATE CONSTRAINT ON (e:Event) ASSERT e.id IS UNIQUE")
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		log.Warn("DDL", "err", err)
	}

	s := &Store{
		db: db,
	}
	return s, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) HasEventHeader(e hash.Event) bool {
	session, err := s.db.Session(neo4j.AccessModeRead)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	id := e.FullID()

	res, err := session.ReadTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		res, err := search(ctx, `MATCH (e:Event {id:'%s'}) RETURN e`, id)
		if err != nil {
			return nil, err
		}

		for res.Next() {
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		panic(err)
	}

	if res.(bool) {
		log.Info("??? has", "id", id)
	}

	return res.(bool)
}

// Load data from events chain.
func (s *Store) Load(events <-chan *inter.Event) error {
	session, err := s.db.Session(neo4j.AccessModeWrite)
	if err != nil {
		return err
	}
	defer session.Close()
	// DML
	var (
		start    = time.Now().Add(-10 * time.Millisecond)
		reported time.Time
		counter  = ratecounter.NewRateCounter(60 * time.Second).WithResolution(1)
		total    int64
		last     string
	)
	for event := range events {
		id := event.Hash().FullID()
		_, err = session.WriteTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
			defer ctx.Close()

			header := Marshal(&event.EventHeaderData)
			log.Debug("<<<", "event", header.String())
			err = exec(ctx, "CREATE (e:Event %s)", header.String())
			if err != nil {
				return nil, err
			}

			for _, p := range event.Parents {
				err = exec(ctx, `MATCH (e:Event {id:'%s'}), (p:Event {id:'%s'}) CREATE (e)-[:PARENT]->(p)`,
					id,
					p.FullID(),
				)
				if err != nil {
					return nil, err
				}
			}

			return nil, ctx.Commit()

		})
		if err != nil {
			log.Error("<<<", "err", err, "event", id)
			// return err
		}

		counter.Incr(1)
		total++
		last = id
		if time.Since(reported) >= statsReportLimit {
			log.Info("<<<", "last", last,
				"per second", counter.Rate()/60,
				"total", total,
				"elapsed", common.PrettyDuration(time.Since(start)))
			reported = time.Now()
		}
	}

	log.Info("Exported events", "last", last,
		"per second", total*1000/time.Since(start).Milliseconds(),
		"total", total,
		"elapsed", common.PrettyDuration(time.Since(start)))
	return nil
}

// FindAncestors of event.
func (s *Store) FindAncestors(event hash.Event) (ancestors []hash.Event, err error) {
	session, err := s.db.Session(neo4j.AccessModeRead)
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
	_, err := ctx.Run(query, nil)
	if err != nil {
		return err
	}

	return nil
}

func search(ctx neo4j.Transaction, cypher string, a ...interface{}) (neo4j.Result, error) {
	query := fmt.Sprintf(cypher, a...)
	log.Debug("cypher", "query", query)
	res, err := ctx.Run(query, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}
