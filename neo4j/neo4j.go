package neo4j

import (
	"fmt"
	"time"

	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru"
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
	db    neo4j.Driver
	cache struct {
		EventsHeaders *lru.Cache
	}
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

	s.cache.EventsHeaders, err = lru.New(500)
	if err != nil {
		panic(err)
	}

	return s, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) HasEventHeader(e hash.Event) bool {
	// Get event from LRU cache first.
	if _, ok := s.cache.EventsHeaders.Get(e); ok {
		return true
	}

	session, err := s.db.Session(neo4j.AccessModeRead)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	id := eventID(e)

	res, err := session.ReadTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		res, err := search(ctx, `MATCH (e:Event %s) RETURN e`, fields{
			"id": id,
		})
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

	return res.(bool)
}

func (s *Store) GetEvent(e hash.Event) *inter.EventHeaderData {
	// Get event from LRU cache first.
	if ev, ok := s.cache.EventsHeaders.Get(e); ok {
		return ev.(*inter.EventHeaderData)
	}

	session, err := s.db.Session(neo4j.AccessModeRead)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	id := eventID(e)

	res, err := session.ReadTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		res, err := search(ctx, `MATCH (e:Event %s) RETURN e.id as id, e.creator as creator`, fields{
			"id": id,
		})
		if err != nil {
			return nil, err
		}

		for res.Next() {
			ff := readFields(res.Record())
			header := new(inter.EventHeaderData)
			unmarshal(ff, header)
			return header, nil
		}
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
	if res == nil {
		return nil
	}
	event := res.(*inter.EventHeaderData)

	res, err = session.ReadTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		res, err := search(ctx, `MATCH (e:Event %s)-[:PARENT]->(p) RETURN p.id`,
			fields{"id": id},
		)
		if err != nil {
			return nil, err
		}
		var parents hash.Events
		for res.Next() {
			p := eventHash(res.Record().GetByIndex(0).(string))
			parents = append(parents, p)
		}
		return parents, nil
	})
	if err != nil {
		panic(err)
	}
	event.Parents = res.(hash.Events)

	return event
}

// Load data from events chain.
func (s *Store) Load(events <-chan *EventData) error {
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
		last     hash.Event
	)
	for edata := range events {
		event := edata.Event
		id := eventID(event.Hash())
		_, err = session.WriteTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
			defer ctx.Close()

			data := marshal(&event.EventHeaderData)
			log.Info("<<<", "event", event.Hash(), "ff", data)
			err = exec(ctx, "CREATE (e:Event %s)", data)
			if err != nil {
				panic(err)
				return nil, err
			}

			for _, p := range event.Parents {
				log.Info("<<<", "event", event.Hash(), "parent", p)
				err = exec(ctx, `MATCH (e:Event %s), (p:Event %s) CREATE (e)-[:PARENT]->(p)`,
					fields{"id": id},
					fields{"id": eventID(p)},
				)
				if err != nil {
					panic(err)
					return nil, err
				}
			}

			return nil, ctx.Commit()

		})
		if err != nil {
			log.Error("<<<", "err", err, "event", event.Hash()) // TODO: why the error is?
			// return err
		}
		s.cache.EventsHeaders.Add(event.Hash(), &event.EventHeaderData)
		if edata.Ready != nil {
			edata.Ready()
		}

		counter.Incr(1)
		total++
		last = event.Hash()
		if time.Since(reported) >= statsReportLimit {
			log.Info("<<<",
				"last", last,
				"per second", counter.Rate()/60,
				"total", total,
				"elapsed", common.PrettyDuration(time.Since(start)))
			reported = time.Now()
		}
	}

	log.Info("Total imported events",
		"last", last,
		"per second", total*1000/time.Since(start).Milliseconds(),
		"total", total,
		"elapsed", common.PrettyDuration(time.Since(start)))
	return nil
}

// FindAncestors of event.
func (s *Store) FindAncestors(e hash.Event) (ancestors []hash.Event, err error) {
	session, err := s.db.Session(neo4j.AccessModeRead)
	if err != nil {
		return
	}
	defer session.Close()

	id := eventID(e)

	res, err := session.ReadTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		res, err := search(ctx, "MATCH (p:Event %s)-[:PARENT*]->(s:Event) RETURN DISTINCT s.id", fields{
			"id": id,
		})
		if err != nil {
			return nil, err
		}

		var ancestors []hash.Event
		for res.Next() {
			pid := eventHash(res.Record().GetByIndex(0).(string))
			ancestors = append(ancestors, pid)
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
