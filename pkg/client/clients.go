package client

import (
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"google.golang.org/grpc"
)

type client struct {
	NodeClient rpc.NodeClient
	Connection *grpc.ClientConn
}

type ClientSet []client

var DefaultClientSet ClientSet

// ConnectToNodes creates the default client set and returns it
func ConnectToNodes(ips []string) (cs ClientSet) {
	for _, target := range ips {
		conn, err := grpc.Dial(target, grpc.WithInsecure())
		if err != nil {
			lg.Log.Warningf("Error during connection setup to node '%s': %s", target, err)
			continue
		}
		cs = append(cs, client{NodeClient: rpc.NewNodeClient(conn), Connection: conn})
		lg.Log.Debugf("Successfully connected to node '%s'", target)
	}

	DefaultClientSet = cs

	return
}

// GetClientSet returns the default client set initialized by ConnectToNodes
func GetClientSet() (cs ClientSet) {
	return DefaultClientSet
}

// ResetBackoff resets the connection backoff for all clients
func (cs ClientSet) ResetBackoff() {
	for _, c := range cs {
		c.Connection.ResetConnectBackoff()
	}
}
