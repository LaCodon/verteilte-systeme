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
	lg.Log.Debugf("Starting new election")

	electing := true
	clientCount := len(config.Default.PeerNodes.Value())
	clients := ConnectToNodes(config.Default.PeerNodes.Value())
	incomingVotes := make(chan *rpc.VoteResponse, clientCount)
	voteCount := 0
	// timeout for RequestVote RPCs
	timeout := time.Duration(1000) * time.Millisecond
	timedOut := make(chan time.Time)

	state.DefaultPersistentState.IncrementCurrentTerm()
	state.DefaultPersistentState.SetCurrentState(state.Candidate)

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
				lg.Log.Warningf("Got error from client.RequestVote call: %s", err)
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
			lg.Log.Debug("Got heartbeat, reverting to follower state")
			state.DefaultPersistentState.SetCurrentState(state.Follower)
			return false
		case <-timedOut:
			lg.Log.Debug("Split votes / timed out waiting for votes")
			// start new election term
			return BeCandidate()
		case resp := <-incomingVotes:
			lg.Log.Debugf("Got vote response: %s", resp)
			if resp.VoteGranted {
				voteCount++

				if voteCount >= clientCount/2 {
					state.DefaultPersistentState.SetCurrentState(state.Leader)
					electing = false
				}
			}
		}
	}

	return true
}
