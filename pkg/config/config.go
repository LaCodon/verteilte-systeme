package config

import (
	"github.com/urfave/cli/v2"
	"time"
)

var Default *Config

type Config struct {
	// PeerNodes holds IPs and ports of the other nodes as strings
	PeerNodes *cli.StringSlice
	// LocalPort holds the listen port of the gRPC server
	LocalPort int
	// HeartbeatInterval is the interval in which the leader will send heartbeats
	HeartbeatInterval time.Duration
	// HeartbeatTimeout is the time after which a follower starts a new election if it doesn't receive a heartbeat from the leader
	HeartbeatTimeout time.Duration
	// RequestVoteTimeout sets the timeout for waiting on RequestVote responses
	RequestVoteTimeout time.Duration
	// AppendEntriesTimeout sets the timeout for waiting on AppendEntries responses
	AppendEntriesTimeout time.Duration
	NodeId           int

	// LogFile holds the filename where the log entries are saved
	Logfile string
}

// RandomizeHeartbeatTimeout sets a new random HeartbeatTimeout
func (c *Config) RandomizeHeartbeatTimeout() {
	c.HeartbeatTimeout = generateRandomHeartbeatTimeout()
}
