package main

import (
	"context"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/client"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/server"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"time"
)

func run(c *cli.Context) error {
	lg.Log.Infof("Hello, I'm node with id %d", config.Default.NodeId)
	lg.Log.Infof("Configured peers: %s", config.Default.PeerNodes)

	// start as follower
	state.DefaultPersistentState.SetCurrentState(state.Follower)

	// the heartbeat channel tracks heartbeats from master node
	client.Heartbeat = make(chan bool, 1)

	if err := server.StartListen(); err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func(c context.Context) {
		// main loop
		for c.Err() == nil {
			switch state.DefaultPersistentState.GetCurrentState() {
			case state.Follower:
				client.BeFollower(c)
			case state.Leader:
				client.BeLeader(c)
			}
		}
	}(ctx)

	go func() {
		// print current state
		for {
			time.Sleep(2000 * time.Millisecond)
			lg.Log.Infof("Current state: %d in term %d", state.DefaultPersistentState.CurrentSate, state.DefaultPersistentState.CurrentTerm)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	// Wait for SIGINT signal
	<-stop

	lg.Log.Info("Stopping service...")

	// Stop routines and RPC server
	cancel()
	// leave some time for go routines to stop
	time.Sleep(1000 * time.Millisecond)
	server.RpcServer.GracefulStop()

	time.Sleep(1 * time.Second)
	return nil
}
