package config

import (
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/urfave/cli/v2"
	"math/rand"
	"time"
)

func init() {
	lg.Log.Debug("Initializing config")
	Default = &Config{
		AllNodes:             cli.NewStringSlice(),
		HeartbeatInterval:    600 * time.Millisecond,
		HeartbeatTimeout:     generateRandomHeartbeatTimeout(),
		RequestVoteTimeout:   500 * time.Millisecond,
		AppendEntriesTimeout: 500 * time.Millisecond,
		UserRequestTimeout:   500 * time.Millisecond,
		RegisterTimeout:      2 * time.Second,
		KickThreshold:        -1,
		LogFormatString: 	  "%d %d %d %s %s\n", // Index Term Action Key Value
	}
}

func generateRandomHeartbeatTimeout() time.Duration {
	// random duration in range 1s - 2s
	return time.Duration(rand.Intn(1000)+1000) * time.Millisecond
}
