package config

import (
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/urfave/cli/v2"
)

func init() {
	lg.Log.Debug("Initializing config")
	Default = &Config{
		PeerNodes: cli.NewStringSlice(),
	}
}
