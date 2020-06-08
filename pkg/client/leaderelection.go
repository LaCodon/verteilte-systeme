package client

import (
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"math/rand"
	"time"
)

// Heartbeat is used to send heartbeats to electionTimeout and thus prevents new leader election
var Heartbeat chan bool

// stopElectionTimeout will stop the electionTimeout gorouting if it gets closed
var stopElectionTimeout chan bool

// BeginElectionTimeout starts watching incoming heartbeats
func BeginElectionTimeout() {
	Heartbeat = make(chan bool, 1)
	stopElectionTimeout = make(chan bool, 1)

	go electionTimeout()
	lg.Log.Info("Started waiting for heartbeats")
}

// EndElectionTimeout ends watching for incoming heartbeats
func EndElectionTimeout() {
	defer func() {
		if r := recover(); r != nil {
			lg.Log.Error("Failed to stop electionTimeout. Was it started before?")
		}
	}()

	close(stopElectionTimeout)
}

// electionTimeout checks if heartbeats are incoming or starts new leader election
func electionTimeout() {
	running := true
	for running {
		// random int between 150 and 300
		timeout := time.Duration(rand.Intn(150) + 150)
		select {
		case <-Heartbeat:
			lg.Log.Debug("Got heartbeat")
		case <-time.After(timeout * time.Millisecond):
			lg.Log.Info("Heartbeat timed out")
			EndElectionTimeout()
			// Todo: start new leader election
		case <-stopElectionTimeout:
			running = false
		}
	}

	lg.Log.Info("Stopped waiting for heartbeats")
}
