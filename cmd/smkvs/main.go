package main

import (
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/urfave/cli/v2"
	"os"
)

var Version = "dev-build"

func main() {
	app := &cli.App{
		Name:        "smKVS",
		Usage:       "Schmidt-Maier-Key-Value-Store",
		Version:     Version,
		Description: "Distributed Key Value Store based on the raft algorithm",
		Action:      run,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:        "peer-ip",
				Aliases:     []string{"i"},
				Usage:       "Peer IPs and Ports in",
				EnvVars:     []string{"SMKVS_PEER_NODES"},
				Required:    true,
				Destination: config.Default.PeerNodes,
			},
			&cli.IntFlag{
				Name:        "local-port",
				Aliases:     []string{"p"},
				Usage:       "RPC server listen port",
				EnvVars:     []string{"SMKVS_PORT"},
				Value:       36000,
				Destination: &config.Default.LocalPort,
			},
		},
		Authors: []*cli.Author{
			{Name: "Fabian Maier", Email: "fabian.maier@fh-erfurt.de"},
			{Name: "Tim Schmidt", Email: "tim.schmidt@fh-erfurt.de"},
		},
	}

	if err := app.Run(os.Args); err != nil {
		lg.Log.Error(err)
		os.Exit(1)
	}
}
