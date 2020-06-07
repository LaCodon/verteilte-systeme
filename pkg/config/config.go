package config

import "github.com/urfave/cli/v2"

var Default *Config

type Config struct {
	// PeerNodes holds IPs and ports of the other nodes as strings
	PeerNodes *cli.StringSlice
	// LocalPort holds the listen port of the gRPC server
	LocalPort int
}
