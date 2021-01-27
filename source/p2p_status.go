package source

import (
	"sync"

	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
)

type status struct {
	sync.RWMutex

	epoch   idx.Epoch
	headers map[hash.Event]bool
}

func newStatus(n idx.Epoch) *status {
	s := new(status)
	s.setEpoch(n)
	return s
}

func (s *status) CurrEpoch() idx.Epoch {
	s.RLock()
	defer s.RUnlock()

	return s.epoch
}

func (s *status) AddHeaders(n idx.Epoch, ee hash.Events) {
	s.Lock()
	defer s.Unlock()

	if n != s.epoch {
		return
	}

	for _, e := range ee {
		if _, set := s.headers[e]; !set {
			s.headers[e] = false
		}
	}
}

func (s *status) IsEpochSealedBy(h hash.Event) bool {
	s.Lock()
	defer s.Unlock()

	if ok, set := s.headers[h]; !set || ok {
		return false
	}

	s.headers[h] = true
	for _, ok := range s.headers {
		if !ok {
			return false
		}
	}

	s.setEpoch(s.epoch + 1)
	return true
}

func (s *status) setEpoch(n idx.Epoch) {
	s.epoch = n
	s.headers = make(map[hash.Event]bool)
}
