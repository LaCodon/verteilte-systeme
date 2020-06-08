package main

import (
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

	if err := server.StartListen(); err != nil {
		return err
	}

	client.BeginElectionTimeout()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	// Wait for SIGINT signal
	<-stop

	lg.Log.Info("Stopping service...")

	// Stop routines
	client.EndHeartbeats()
	server.RpcServer.GracefulStop()

	time.Sleep(1 * time.Second)
	return nil
}
