package client

import (
	"context"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"time"
)

func BeLeader(ctx context.Context) {
	for state.DefaultPersistentState.GetCurrentState() == state.Leader && ctx.Err() == nil {
		term := state.DefaultPersistentState.GetCurrentTerm()

		for _, client := range DefaultClientSet {
			go func(c rpc.NodeClient, ctx context.Context) {
				ctx, _ = context.WithTimeout(ctx, config.Default.AppendEntriesTimeout)
				_, err := c.AppendEntries(ctx, &rpc.AppendEntriesRequest{Term: term})
				if err != nil {
					lg.Log.Debugf("error from client.AppendEntries: %s", err)
				}
			}(client.NodeClient, ctx)
		}

		time.Sleep(config.Default.HeartbeatInterval)
	}
}
