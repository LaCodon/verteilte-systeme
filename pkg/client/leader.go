package client

import (
	"context"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
)

func BeLeader() {
	clients := ConnectToNodes(config.Default.PeerNodes.Value())
	for {
		term := state.DefaultPersistentState.GetCurrentTerm()

		for _, c := range clients {
			_, err := c.AppendEntries(context.Background(), &rpc.AppendEntriesRequest{Term: term})
			if err != nil {
				lg.Log.Errorf("error from client.AppendEntries: %s", err)
			}
		}

	}
}
