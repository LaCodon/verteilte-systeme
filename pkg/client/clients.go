package client

import (
	"github.com/LaCodon/verteilte-systeme/internal/helper"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"google.golang.org/grpc"
	"sync"
)

type Client struct {
	NodeClient rpc.NodeClient
	Connection *grpc.ClientConn
	Target     string
	ErrorCount int
}
type ClientSet []*Client

var ForceClientReconnect bool

var defaultClientSet ClientSet
var clientConnectMutex sync.Mutex

// connectToNodes creates the default Client set and returns it
func ConnectToNodes(ips []string) (cs ClientSet) {
	for _, c := range defaultClientSet {
		if err := c.Connection.Close(); err != nil {
			lg.Log.Warningf("Failed to close connection before reconnect")
		}
	}

	for _, target := range ips {
		nc, err := NewClient(target)
		if err != nil {
			lg.Log.Warningf("Error during connection setup to node '%s': %s", target, err)
			continue
		}

		cs = append(cs, nc)
		lg.Log.Debugf("Successfully connected to node '%s'", target)
	}

	defaultClientSet = cs

	return
}

// NewClient calls grpc.Dial and returns the client object
func NewClient(target string) (*Client, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{NodeClient: rpc.NewNodeClient(conn), Connection: conn, Target: target}, nil
}

// GetClientSet returns the default Client set initialized by connectToNodes
func GetClientSet() (cs ClientSet) {
	clientConnectMutex.Lock()
	defer clientConnectMutex.Unlock()

	if ForceClientReconnect {
		lg.Log.Infof("Force connection reestablishment")
		ConnectToNodes(config.Default.GetPeerNodesData())
		ForceClientReconnect = false
	}

	return defaultClientSet
}

// ResetBackoff resets the connection backoff for all clients
func (cs ClientSet) ResetBackoff() {
	for _, c := range cs {
		c.Connection.ResetConnectBackoff()
	}
}

// HasClient returns true if the client set has a node with nodeId
func (cs ClientSet) HasClient(nodeId uint32) bool {
	for _, c := range cs {
		if c.GetId() == nodeId {
			return true
		}
	}

	return false
}

// GetId hashes the target and thus makes it a unique id
func (c *Client) GetId() uint32 {
	return helper.TargetToId(c.Target)
}
