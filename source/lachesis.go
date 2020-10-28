package source

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/Fantom-foundation/go-lachesis/app"
	"github.com/Fantom-foundation/go-lachesis/gossip"
	"github.com/Fantom-foundation/go-lachesis/integration"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/kvdb/flushable"
	"github.com/ethereum/go-ethereum/log"
)

func Events(ctx context.Context, dataDir string, from, to idx.Epoch) <-chan *inter.Event {
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
