package main

import (
	"testing"
)

func TestTransfersGenerator(t *testing.T) {
	t.Skip("example only")
	cfg := &Config{
		ChainId: 999,
		Accs: struct {
			Count  uint
			Offset uint
		}{
			Count:  10,
			Offset: 100,
		},
	}
	g := NewTransfersGenerator(cfg, 1, 1)
	for i := 0; i < 2*len(g.accs); i++ {
		_ = g.Yield()
	}
}
