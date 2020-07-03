package main

import (
	"context"
	"fmt"
	"github.com/LaCodon/verteilte-systeme/internal/helper"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/client"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"github.com/LaCodon/verteilte-systeme/pkg/server"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"sync"
	"time"
)

func run(c *cli.Context) error {
	config.Default.NodeId = helper.TargetToId(config.Default.MyNode)

	lg.Log.Infof("Hello, I'm node with id %d", config.Default.NodeId)
	lg.Log.Infof("Configured peers: %s", config.Default.GetPeerNodesData())

	// initialize list of all known nodes for propagation
	allNodes := append(config.Default.AllNodes.Value(), config.Default.MyNode)
	config.Default.SetNewNodes(allNodes)

	// check config
	if config.Default.HeartbeatInterval < config.Default.AppendEntriesTimeout {
		return fmt.Errorf("heartbeat intveral has to be greater than append entries timeout")
	}

	config.Default.Logfile = fmt.Sprintf("log%d.txt", config.Default.NodeId)

	// start as follower
	state.DefaultPersistentState.SetCurrentState(state.Follower)

	// the heartbeat channel tracks heartbeats from master node
	client.Heartbeat = make(chan bool, 1)

	// initial connect to other nodes
	client.ForceClientReconnect = true
	client.GetClientSet()

	ctx, cancel := context.WithCancel(context.Background())

	// add this node to the cluster
	register(ctx)

	if config.Default.PeerNodeCount() == 0 {
		return fmt.Errorf("registration failed, no leader answered")
	}

	// start listening for incoming requests
	if err := server.StartListen(); err != nil {
		return err
	}

	// main state machine
	go func(c context.Context) {
		// main loop
		for c.Err() == nil {
			switch state.DefaultPersistentState.GetCurrentState() {
			case state.Follower:
				client.BeFollower(c)
			case state.Leader:
				client.BeLeader(c)
			}
		}
	}(ctx)

	// debugging help
	go func() {
		// print current state
		for {
			time.Sleep(2000 * time.Millisecond)
			lg.Log.Infof("Current state: %d in term %d", state.DefaultPersistentState.CurrentSate, state.DefaultPersistentState.CurrentTerm)
			lg.Log.Infof("Current peer nodes: %s", config.Default.GetPeerNodesData())
			//lg.Log.Warningf("Current Log: %s", state.DefaultPersistentState.Log)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	// Wait for SIGINT signal
	<-stop

	lg.Log.Info("Stopping service...")

	// Stop routines and RPC server
	cancel()
	// leave some time for go routines to stop
	time.Sleep(1000 * time.Millisecond)
	server.RpcServer.GracefulStop()

	time.Sleep(1 * time.Second)
	return nil
}

// register adds this node to the cluster
func register(ctx context.Context) {
	lg.Log.Infof("Registering node...")
	var wg sync.WaitGroup

	for _, other := range client.GetClientSet() {
		wg.Add(1)
		go func(cl *client.Client) {
			defer wg.Done()

			c, _ := context.WithTimeout(ctx, config.Default.RegisterTimeout)
			resp, err := cl.NodeClient.RegisterNode(c, &rpc.NodeRegisterRequest{
				ConnectionData: config.Default.MyNode,
				NodeId:         config.Default.NodeId,
			})
			if err != nil {
				lg.Log.Errorf("Error during registration process: %s", err.Error())
				return
			}

			if resp.Success == false && resp.RedirectTarget != "" {
				// remove non leaders from node list but only if they sent an alternative where to register
				config.Default.RemoveNode(cl.Target)
				client.ForceClientReconnect = true
			}
		}(other)
	}

	wg.Wait()
	lg.Log.Infof("Registering node done.")
}
