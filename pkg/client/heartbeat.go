package client

import (
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"time"
)

var stopHeartbeats chan bool

// BeginHeartbeats starts sending AppendEntries RPCs. Call this as soon as node is leader
func BeginHeartbeats() {
	stopHeartbeats = make(chan bool, 1)
	go doHeartbeats()
}

// EndHeartbeats stops sending AppendEntries RPCs
func EndHeartbeats() {
	defer func() {
		if r := recover(); r != nil {
			lg.Log.Error("Failed to stop heartbeating. Was it started before?")
		}
	}()

	close(stopHeartbeats)
}

// doHeartbeats is used as goroutine to send periodical heartbeats
func doHeartbeats() {
	running := true
	for running {
		select {
		case <-stopHeartbeats:
			running = false
		case <-time.After(config.Default.HeartbeatTimeout):
			lg.Log.Debug("Heartbeat triggered by HeartbeatTimeout")
			// TODO: issue concurrent AppendEntries RPCs to followers
		}
	}

	lg.Log.Info("Stopped heartbeating")
}
