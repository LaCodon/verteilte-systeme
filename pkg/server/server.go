package server

import (
	"fmt"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"google.golang.org/grpc"
	"net"
)

// RpcServer holds a reference to the server after it has been started with StartListen
var RpcServer *grpc.Server

type Server struct {
	rpc.NodeServer
}

// StartListen start the RPC server in a non blocking manner
func StartListen() error {
	iface := fmt.Sprintf("0.0.0.0:%d", config.Default.LocalPort)
	listener, err := net.Listen("tcp", iface)
	if err != nil {
		return err
	}

	lg.Log.Infof("Listening on %s", iface)

	// TODO: secure connections via TLS and add client authentication
	RpcServer = grpc.NewServer()
	rpc.RegisterNodeServer(RpcServer, &Server{})

	go func() {
		lg.Log.Info("RPC Server is now serving")
		if err := RpcServer.Serve(listener); err != nil {
			lg.Log.Criticalf("Failed to start RPC server: %RpcServer", err.Error())
		}
	}()

	return nil
}
