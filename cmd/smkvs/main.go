package main

import (
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/urfave/cli/v2"
	"math/rand"
	"os"
	"time"
)

var Version = "dev-build"

func main() {
	rand.Seed(time.Now().UnixNano())

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
				Destination: config.Default.AllNodes,
			},
			&cli.IntFlag{
				Name:        "local-port",
				Aliases:     []string{"p"},
				Usage:       "RPC server listen port",
				EnvVars:     []string{"SMKVS_PORT"},
				Value:       36000,
				Destination: &config.Default.LocalPort,
			},
			&cli.IntFlag{
				Name:        "node-id",
				Aliases:     []string{"n"},
				Usage:       "ID of the current node",
				EnvVars:     []string{"SMKVS_NODE_ID"},
				Value:       rand.Int(),
				Destination: &config.Default.NodeId,
			},
			&cli.StringFlag{
				Name:        "listen",
				Aliases:     []string{"l"},
				Usage:       "External listener IP and Port for current node",
				EnvVars:     []string{"SMKVS_NODE_LISTEN"},
				Required:    true,
				Value:       "127.0.0.1:36000",
				Destination: &config.Default.MyNode,
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
