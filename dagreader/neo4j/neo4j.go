package neo4j

import (
	"fmt"
	"sync"
	"time"

	"github.com/Fantom-foundation/go-opera/logger"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/paulbellamy/ratecounter"

	"github.com/Fantom-foundation/lachesis-dag-tool/dagreader/internal"
)

const (
	DefaultDb = "bolt://localhost:7687"

	// statsReportLimit is the time limit during import and export after which we
	// always print out progress. This avoids the user wondering what's going on.
	statsReportLimit = 8 * time.Second
)

type Db struct {
	drv   neo4j.Driver
	busy  sync.WaitGroup
	cache struct {
		EventInfos *lru.Cache
	}

	logger.Instance
}

func New(dbUrl string) (*Db, error) {
	db, err := neo4j.NewDriver(dbUrl, neo4j.NoAuth(), func(c *neo4j.Config) {
		c.Encrypted = false
	})
	if err != nil {
		return nil, err
	}

	s := &Db{
		drv:      db,
		Instance: logger.MakeInstance(),
	}
	s.SetName("neo4j")

	s.busy.Add(1)
	defer s.busy.Done()

	session, err := db.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	DDLs := []string{
		"CREATE CONSTRAINT ON (e:Event) ASSERT e.id IS UNIQUE",
		"CREATE CONSTRAINT ON (b:Block) ASSERT b.id IS UNIQUE",
		"CREATE (b:LastBlock {id:'current',num:2})",
	}
	for _, query := range DDLs {
		_, err = session.WriteTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
			defer ctx.Close()

			err := exec(ctx, query)
			if err != nil {
				log.Warn("DDL", "err", err, "query", query)
				return nil, err
			}

			return nil, ctx.Commit()
		})
		if err != nil {
			ignoreFakeError(err)
		}
	}

	s.cache.EventInfos, err = lru.New(500)
	if err != nil {
		panic(err)
	}

	return s, nil
}

func (s *Db) Close() error {
	s.busy.Wait()
	return s.drv.Close()
}

func (s *Db) HasEvent(e hash.Event) bool {
	// Get event from LRU cache first.
	if _, ok := s.cache.EventInfos.Get(e); ok {
		return true
	}

	s.busy.Add(1)
	defer s.busy.Done()

	session, err := s.drv.Session(neo4j.AccessModeRead)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	res, err := session.ReadTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		cursor, err := search(ctx, `MATCH (e:Event %s) RETURN e`, fields{
			"id": eventID(e),
		})
		if err != nil {
			panic(err)
		}

		has := cursor.Next()
		return has, nil
	})
	if err != nil {
		ignoreFakeError(err)
	}

	return res.(bool)
}

// GetEvent returns event info.
func (s *Db) GetEvent(e hash.Event) *internal.EventInfo {
	// Get event from LRU cache first.
	if ev, ok := s.cache.EventInfos.Get(e); ok {
		return ev.(*internal.EventInfo)
	}

	s.busy.Add(1)
	defer s.busy.Done()

	session, err := s.drv.Session(neo4j.AccessModeRead)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	res, err := session.ReadTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		cursor, err := search(ctx, `MATCH (e:Event %s) RETURN e.block as block, e.role as role, e.id as id, e.creator as creator`, fields{
			"id": eventID(e),
		})
		if err != nil {
			panic(err)
		}

		for cursor.Next() {
			ff := readFields(cursor.Record())
			return ff, nil
		}
		return nil, nil
	})
	if err != nil {
		ignoreFakeError(err)
	}
	if res == nil {
		return nil
	}

	ff := res.(fields)
	ff["parents"] = s.getParents(session, e)

	info := new(internal.EventInfo)
	unmarshal(ff, info)

	return info
}

func (s *Db) getParents(session neo4j.Session, e hash.Event) hash.Events {
	var parents hash.Events
	id := eventID(e)
	_, err := session.ReadTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		cursor, err := search(ctx, `MATCH (e:Event %s)-[:PARENT]->(p) RETURN p.id`,
			fields{"id": id},
		)
		if err != nil {
			panic(err)
		}
		for cursor.Next() {
			p := eventHash(cursor.Record().GetByIndex(0).(string))
			parents = append(parents, p)
		}
		return nil, nil
	})
	if err != nil {
		ignoreFakeError(err)
	}

	return parents
}

// Load data from input chain.
func (s *Db) Load(events <-chan *internal.EventInfo) {
	s.busy.Add(1)
	defer s.busy.Done()

	session, err := s.drv.Session(neo4j.AccessModeWrite)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	parents := make(chan *internal.EventInfo, 10)
	defer close(parents)
	go s.loadParents(parents)

	var lastBlock idx.Block = s.GetLastBlock()
	for info := range events {
		_, err = session.WriteTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
			defer ctx.Close()

			if lastBlock < info.Block {
				lastBlock = info.Block
				s.setLastBlock(lastBlock)
			}

			data := marshal(info)
			delete(data, "parents")
			s.Log.Info("<<< event", "id", info.Event.ID(), "data", data)
			err = exec(ctx, "CREATE (e:Event %s)", data)
			if err != nil {
				panic(err)
			}

			return nil, ctx.Commit()
		})
		if err != nil {
			ignoreFakeError(err)
		}

		parents <- info
	}
}

func (s *Db) loadParents(events <-chan *internal.EventInfo) {
	s.busy.Add(1)
	defer s.busy.Done()

	session, err := s.drv.Session(neo4j.AccessModeWrite)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var (
		start    = time.Now().Add(-10 * time.Millisecond)
		reported time.Time
		counter  = ratecounter.NewRateCounter(60 * time.Second).WithResolution(1)
		total    int64
		last     hash.Event
	)

	for info := range events {
		event := info.Event
		id := event.ID()
		_, err = session.WriteTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
			defer ctx.Close()

			for _, p := range event.Parents() {
				pid := eventID(p)
				err = exec(ctx, `MATCH (e:Event %s), (p:Event %s) CREATE (e)-[:PARENT]->(p)`,
					fields{"id": eventID(id)},
					fields{"id": pid},
				)
				if err != nil {
					panic(err)
				}
			}
			return nil, ctx.Commit()
		})
		if err != nil {
			ignoreFakeError(err)
		}

		s.cache.EventInfos.Add(id, info)
		info.Done()

		counter.Incr(1)
		total++
		last = event.ID()
		if time.Since(reported) >= statsReportLimit {
			s.Log.Info("<<<",
				"last", last,
				"rate", counter.Rate()/60,
				"total", total,
				"elapsed", common.PrettyDuration(time.Since(start)))
			reported = time.Now()
		}
	}

	s.Log.Info("Total imported events",
		"last", last,
		"rate", total*1000/time.Since(start).Milliseconds(),
		"total", total,
		"elapsed", common.PrettyDuration(time.Since(start)))
}

// FindAncestors of event.
func (s *Db) FindAncestors(e hash.Event) []hash.Event {
	s.busy.Add(1)
	defer s.busy.Done()

	session, err := s.drv.Session(neo4j.AccessModeRead)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	id := eventID(e)

	res, err := session.ReadTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		cursor, err := search(ctx, "MATCH (p:Event %s)-[:PARENT*]->(s:Event) RETURN DISTINCT s.id", fields{
			"id": id,
		})
		if err != nil {
			panic(err)
		}

		var ancestors []hash.Event
		for cursor.Next() {
			pid := eventHash(cursor.Record().GetByIndex(0).(string))
			ancestors = append(ancestors, pid)
		}
		return ancestors, nil
	})
	if err != nil {
		ignoreFakeError(err)
	}

	return res.([]hash.Event)
}

func (s *Db) setLastBlock(num idx.Block) {
	s.busy.Add(1)
	defer s.busy.Done()

	session, err := s.drv.Session(neo4j.AccessModeWrite)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	const key = "current"

	_, err = session.WriteTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		defer ctx.Close()

		err := exec(ctx, `MATCH (e:LastBlock %s) SET e.num = %d`,
			fields{"id": key}, num)
		if err != nil {
			panic(err)
		}

		return nil, ctx.Commit()
	})
	if err != nil {
		ignoreFakeError(err)
	}
}

func (s *Db) GetLastBlock() idx.Block {
	s.busy.Add(1)
	defer s.busy.Done()

	session, err := s.drv.Session(neo4j.AccessModeRead)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	const key = "current"

	res, err := session.ReadTransaction(func(ctx neo4j.Transaction) (interface{}, error) {
		cursor, err := search(ctx, `MATCH (e:LastBlock %s) RETURN e.num as num`, fields{
			"id": key,
		})
		if err != nil {
			panic(err)
		}

		for cursor.Next() {
			b := idx.Block(cursor.Record().GetByIndex(0).(int64))
			return b, nil
		}
		return nil, nil
	})
	if err != nil {
		ignoreFakeError(err)
	}
	if res == nil {
		return idx.Block(2)
	}
	return res.(idx.Block)
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

func ignoreFakeError(err error) {
	log.Trace("neo4j non critical error", "err", err)
}
