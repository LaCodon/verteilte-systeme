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
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "Run a smkvs node",
				Action: run,
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
			},
			{
				Name:   "set",
				Usage:  "Store a new key value pair",
				Action: setKVPair,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "ip",
						Aliases:     []string{"i"},
						Usage:       "IP of a storage node, where the key value pair should be stored",
						Value:       "127.0.0.1:36000",
						Destination: &nodeIp,
					},
					&cli.StringFlag{
						Name:        "key",
						Aliases:     []string{"k"},
						Usage:       "Value is stored to this key",
						Required:    true,
						Destination: &userInput.Key,
					},
					&cli.StringFlag{
						Name:        "value",
						Aliases:     []string{"v"},
						Usage:       "Value to be stored",
						Required:    true,
						Destination: &userInput.Var,
					},
				},
			},
			{
				Name:   "delete",
				Usage:  "Delete a stored key value pair",
				Action: deleteKVPair,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "ip",
						Aliases:     []string{"i"},
						Usage:       "IP of a storage node, where the key value pair should be stored",
						Value:       "127.0.0.1:36000",
						Destination: &nodeIp,
					},
					&cli.StringFlag{
						Name:        "key",
						Aliases:     []string{"k"},
						Usage:       "Value is stored to this key",
						Required:    true,
						Destination: &userInput.Key,
					},
				},
			},
			{
				Name:   "get",
				Usage:  "Get all current key value pairs stored in the system.",
				Action: getStorage,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "ip",
						Aliases:     []string{"i"},
						Usage:       "IP of a storage node, where the key value pairs are stored",
						Value:       "127.0.0.1:36000",
						Destination: &nodeIp,
					},
					&cli.StringFlag{
						Name:        "key",
						Aliases:     []string{"k"},
						Usage:       "Returns only the value to the specified key",
						Destination: &userInput.Key,
					},
				},
			},
		},
		Authors: []*cli.Author{
			{Name: "Fabian Maier", Email: "fabian.maier@fh-erfurt.de"},
			{Name: "Tim Schmidt", Email: "tim.schmidt@fh-erfurt.de"},
			{Name: "Tim Schmidt", Email: "tim.schmidt.1@fh-erfurt.de"},
		},
	}

	if err := app.Run(os.Args); err != nil {
		lg.Log.Error(err)
		os.Exit(1)
	}
}
