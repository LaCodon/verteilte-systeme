package main

import (
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/server"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
)

func run(c *cli.Context) error {
	lg.Log.Infof("Configured peers: %s", config.Default.PeerNodes)

	if err := server.StartListen(); err != nil {
		return err
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	// Wait for SIGINT signal
	<-stop

	lg.Log.Info("Stopping service...")

	// Stop routines
	server.RpcServer.GracefulStop()

	return nil
}
