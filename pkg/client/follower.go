package client

import (
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"math/rand"
	"time"
)

// Heartbeat is used to send heartbeats to electionTimeout and thus prevents new leader election
var Heartbeat chan bool

// BeFollower starts watching incoming heartbeats
func BeFollower() {
	lg.Log.Info("Started waiting for heartbeats")

	running := true
	for running {
		// random int between 150 and 300
		timeout := time.Duration(rand.Intn(1000)+1000) * time.Millisecond
		select {
		case <-Heartbeat:
			lg.Log.Debug("Got heartbeat")
		case <-time.After(timeout):
			lg.Log.Info("Heartbeat timed out")
			if BeCandidate() {
				running = false
				lg.Log.Debug("I'm master now!")
			}
		}
	}

	lg.Log.Info("Stopped waiting for heartbeats")
}
