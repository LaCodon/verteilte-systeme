package server

import (
	"context"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/client"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
)

func (s *Server) RegisterNode(c context.Context, ar *rpc.NodeRegisterRequest) (*rpc.NodeRegisterResponse, error) {
	if state.DefaultPersistentState.CurrentSate != state.Leader {
		return &rpc.NodeRegisterResponse{Success: false, RedirectTarget: state.DefaultVolatileState.GetCurrentLeader()}, nil
	}

	for _, v := range config.Default.AllNodes.Value() {
		// don't add a node twice
		if v == ar.ConnectionData {
			lg.Log.Debugf("Skipped adding node because already knew it")
			return &rpc.NodeRegisterResponse{Success: true}, nil
		}
	}

	config.Default.AddNode(ar.ConnectionData)
	client.ForceClientReconnect = true

	lg.Log.Infof("Successfully registered new node %s, %s", ar.ConnectionData, config.Default.AllNodes.Value())

	return &rpc.NodeRegisterResponse{Success: true}, nil
}
