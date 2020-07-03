package config

import (
	"github.com/urfave/cli/v2"
	"sync"
	"time"
)

var Default *Config

var allNodeLock sync.Mutex

type Config struct {
	// AllNodes holds IPs and ports of the ALL nodes as strings
	AllNodes *cli.StringSlice
	// MyNode holds own connection string (target)
	MyNode string
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
	// RegisterTimeout is the timeout for RegisterNode calls
	RegisterTimeout time.Duration
	// KickThreshold is the max amount of failed AppendEntries RPCs after a node gets kicked from the leader
	KickThreshold int
	NodeId        uint32

	// LogFile holds the filename where the log entries are saved
	Logfile string
}

// RandomizeHeartbeatTimeout sets a new random HeartbeatTimeout
func (c *Config) RandomizeHeartbeatTimeout() {
	c.HeartbeatTimeout = generateRandomHeartbeatTimeout()
}

// SetNewNodes exchanges current node list with targets
func (c *Config) SetNewNodes(targets []string) {
	allNodeLock.Lock()
	defer allNodeLock.Unlock()

	c.AllNodes = cli.NewStringSlice(targets...)
}

// AddNode adds a node to the current node list
func (c *Config) AddNode(target string) {
	allNodeLock.Lock()
	defer allNodeLock.Unlock()

	nodes := c.AllNodes.Value()
	nodes = append(nodes, target)

	c.AllNodes = cli.NewStringSlice(nodes...)
}

// RemoveNode deletes a node from the current node list
func (c *Config) RemoveNode(target string) {
	allNodeLock.Lock()
	defer allNodeLock.Unlock()

	nodes := make([]string, 0, len(c.AllNodes.Value()))

	for _, n := range c.AllNodes.Value() {
		if n != target {
			nodes = append(nodes, n)
		}
	}

	c.AllNodes = cli.NewStringSlice(nodes...)
}

func (c *Config) HasNodesDiff(compare []string) bool {
	allNodeLock.Lock()
	defer allNodeLock.Unlock()

	current := c.AllNodes.Value()

	if len(compare) != len(current) {
		return true
	}

	hasCount := 0
	for _, x := range current {
		for _, y := range compare {
			if x == y {
				hasCount++
				break
			}
		}
	}

	return hasCount == len(current)
}

// PeerNodeCount returns amount of peer nodes (this node exclusive)
func (c *Config) PeerNodeCount() int {
	allNodeLock.Lock()
	defer allNodeLock.Unlock()

	count := len(c.AllNodes.Value())

	// AllNodes may contain this node -> subtract it
	for _, n := range c.AllNodes.Value() {
		if n == c.MyNode {
			count--
			break
		}
	}

	return count
}

// GetPeerNodesData returns IPs and ports of all currently known peer nodes
func (c *Config) GetPeerNodesData() []string {
	allNodeLock.Lock()
	defer allNodeLock.Unlock()

	// don't add myself do return slice
	peerNodes := make([]string, 0, len(c.AllNodes.Value()))
	for _, n := range c.AllNodes.Value() {
		if n != c.MyNode {
			peerNodes = append(peerNodes, n)
		}
	}

	return peerNodes
}
