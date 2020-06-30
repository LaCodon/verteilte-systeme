package client

import (
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"google.golang.org/grpc"
)

type ClientSet []rpc.NodeClient

var connections ClientSet

func ConnectToNodes(ips []string) (cs ClientSet) {
	if len(connections) == 0 {
		for _, target := range ips {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				lg.Log.Warningf("Error during connection setup to node '%s': %s", target, err)
				continue
			}
			cs = append(cs, rpc.NewNodeClient(conn))
			lg.Log.Debugf("Successfully connected to node '%s'", target)
		}
	}
	return
}
