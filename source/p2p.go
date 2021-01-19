package source

import (
	"context"

	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
)

func EventsFromP2p(ctx context.Context, network string, from, to idx.Epoch) <-chan *inter.Event {
	log.Info("Events of epoches", "from", from, "to", to, "network", network)
	output := make(chan *inter.Event, 10)

	go func() {
		defer close(output)

		stack := makeStack(output)

		err := stack.Start()
		if err != nil {
			log.Error("Error starting protocol stack", "err", err)
			return
		}
		defer stack.Close()

		<-ctx.Done()
	}()

	return output
}

func makeStack(chan<- *inter.Event) *node.Node {
	// TODO: make it
	return nil
}
