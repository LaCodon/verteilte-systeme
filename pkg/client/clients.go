package client

import (
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"google.golang.org/grpc"
)

type ClientSet []rpc.NodeClient

func ConnectToNodes(ips []string) (cs ClientSet) {
	for _, target := range ips {
		conn, err := grpc.Dial(target, grpc.WithInsecure())
		if err != nil {
			lg.Log.Warningf("Error during connection setup to node '%s': %s", target, err)
			continue
		}
		cs = append(cs, rpc.NewNodeClient(conn))
		lg.Log.Info("Successfully connected to node '%s'", target)
	}

	return
}
