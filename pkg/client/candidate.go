package client

import (
	"context"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"time"
)

// elect requests votes from other clients to make this node the new leader.
// Returns true if this node has been elected
func BeCandidate() bool {
	lg.Log.Info("Starting new election")

	{ // transition to candidate state
		state.DefaultPersistentState.Mutex.Lock()

		id := int32(config.Default.NodeId)
		state.DefaultPersistentState.UpdateFragile(state.Candidate, state.DefaultPersistentState.CurrentTerm+1, &id)

		state.DefaultPersistentState.Mutex.Unlock()
	}

	electing := true
	nodeCount := len(config.Default.PeerNodes.Value()) + 1
	clients := ConnectToNodes(config.Default.PeerNodes.Value())
	incomingVotes := make(chan *rpc.VoteResponse, nodeCount)
	// start with voteCount = 1 because this node votes for itself
	voteCount := 1
	// timeout for RequestVote RPCs
	timeout := time.Duration(1000) * time.Millisecond
	timedOut := make(chan time.Time)

	state.DefaultPersistentState.Mutex.RLock()
	r := &rpc.VoteRequest{
		Term:         state.DefaultPersistentState.CurrentTerm,
		CandidateId:  int32(config.Default.NodeId),
		LastLogIndex: state.DefaultPersistentState.GetLastLogIndexFragile(),
		LastLogTerm:  state.DefaultPersistentState.GetLastLogTermFragile(),
	}
	state.DefaultPersistentState.Mutex.RUnlock()

	lg.Log.Debugf("Requesting vote for term %d, candidID %d, lIndex %d, lTerm %d", r.Term, r.CandidateId, r.LastLogIndex, r.LastLogTerm)

	// send vote requests to all other nodes
	for _, c := range clients {
		go func(client rpc.NodeClient) {
			ctx, _ := context.WithTimeout(context.Background(), timeout)
			resp, err := client.RequestVote(ctx, r)
			if err == nil {
				if resp.VoteGranted {
					incomingVotes <- resp
				}
			} else {
				lg.Log.Debugf("Got error from client.RequestVote call: %s", err)
			}
		}(c)
	}

	go func() {
		t := <-time.After(timeout)
		timedOut <- t
	}()

	for electing {
		select {
		case <-Heartbeat:
			lg.Log.Info("Got heartbeat, reverting to follower state")
			state.DefaultPersistentState.SetCurrentState(state.Follower)
			return false
		case <-timedOut:
			lg.Log.Info("Split votes / timed out waiting for votes")
			state.DefaultPersistentState.SetCurrentState(state.Follower)
			return false
		case resp := <-incomingVotes:
			lg.Log.Infof("Got vote response: %s", resp)
			if resp.VoteGranted {
				voteCount++

				if voteCount > nodeCount/2 {
					state.DefaultPersistentState.SetCurrentState(state.Leader)
					electing = false
				}
			}
		}
	}

	return true
}
