package client

import (
	"context"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"time"
)

// Heartbeat is used to send heartbeats to electionTimeout and thus prevents new leader election
var Heartbeat chan bool

// BeFollower starts watching incoming heartbeats
func BeFollower(ctx context.Context) {
	lg.Log.Info("Started waiting for heartbeats")

	running := true
	for running && ctx.Err() == nil {
		// Randomize timeout every loop to prevent from forcing a node into special role
		config.Default.RandomizeHeartbeatTimeout()
		select {
		case <-Heartbeat:
			lg.Log.Debug("Got heartbeat")
		case <-time.After(config.Default.HeartbeatTimeout):
			lg.Log.Info("Heartbeat timed out")
			if BeCandidate() {
				state.DefaultVolatileState.CurrentLeader = ""
				running = false
				lg.Log.Infof("I'm master now for term %d", state.DefaultPersistentState.GetCurrentTerm())
			}
		}
	}

	lg.Log.Info("Stopped waiting for heartbeats")
}
