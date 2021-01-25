package source

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/Fantom-foundation/go-lachesis/app"
	"github.com/Fantom-foundation/go-lachesis/gossip"
	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/integration"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/kvdb/flushable"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Fantom-foundation/lachesis-dag-tool/neo4j"
)

func EventsFromDatadir(ctx context.Context, dataDir string, from, to idx.Epoch, store *neo4j.Store) <-chan *inter.Event {
	log.Info("Events of epoches", "from", from, "to", to, "datadir", dataDir)
	output := make(chan *inter.Event, 10)

	go func() {
		defer close(output)

		gdb := makeGossipStore(dataDir)
		defer gdb.Close()

		gdb.ForEachEvent(from, func(event *inter.Event) bool {
			if to >= from && event.Epoch > to {
				return false
			}

			if store.HasEventHeader(event.Hash()) {
				return true
			}

			select {
			case <-ctx.Done():
				return false
			case output <- event:
				log.Debug(">>>", "event", event.Hash())
			}
			return true
		})
	}()

	return output
}

func makeGossipStore(dataDir string) *gossip.Store {
	dbs := flushable.NewSyncedPool(integration.DBProducer(dataDir))
	gdb := gossip.NewStore(dbs, gossip.LiteStoreConfig(), app.LiteStoreConfig())
	gdb.SetName("lachesis-db")
	return gdb
}

// FindAncestors of event.
func FindAncestors(dataDir string, event hash.Event) (ancestors []hash.Event, err error) {
	const (
		processed = true
	)

	gdb := makeGossipStore(dataDir)
	defer gdb.Close()

	e0 := gdb.GetEvent(event)
	if e0 == nil {
		return
	}
	queue := make(map[hash.Event]bool)
	for _, p := range e0.Parents {
		queue[p] = !processed
	}

	repeat := true
	for repeat {
		repeat = false
		for h, status := range queue {
			if status == processed {
				continue
			}

			queue[h] = processed
			e := gdb.GetEvent(h)
			for _, p := range e.Parents {
				if _, was := queue[p]; !was {
					queue[p] = !processed
				}
			}

			repeat = true
			break
		}
	}

	for p := range queue {
		ancestors = append(ancestors, p)
	}

	return
}

// DefaultDataDir is the default data directory to use for the databases and other
// persistence requirements.
func DefaultDataDir() string {
	// Try to place the data folder in the user's home dir
	home := homeDir()
	if home != "" {
		switch runtime.GOOS {
		case "darwin":
			return filepath.Join(home, "Library", "Lachesis")
		case "windows":
			// We used to put everything in %HOME%\AppData\Roaming, but this caused
			// problems with non-typical setups. If this fallback location exists and
			// is non-empty, use it, otherwise DTRT and check %LOCALAPPDATA%.
			fallback := filepath.Join(home, "AppData", "Roaming", "Lachesis")
			appdata := windowsAppData()
			if appdata == "" || isNonEmptyDir(fallback) {
				return fallback
			}
			return filepath.Join(appdata, "Lachesis")
		default:
			return filepath.Join(home, ".lachesis")
		}
	}
	// As we cannot guess a stable location, return empty and handle later
	return ""
}
