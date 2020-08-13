package server

import (
	"context"
	"fmt"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/client"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
)

// RequestVote gets called by a candidate for leader election
func (s *Server) RequestVote(c context.Context, v *rpc.VoteRequest) (*rpc.VoteResponse, error) {
	lg.Log.Debugf("Got vote request from %d with term %d", v.CandidateId, v.Term)

	if !client.GetClientSet().HasClient(v.CandidateId) {
		lg.Log.Debugf("Rejected vote request from %d because it's not part of known cluster", v.CandidateId)
		return &rpc.VoteResponse{
			Term:        state.DefaultPersistentState.GetCurrentTerm(),
			VoteGranted: false,
		}, fmt.Errorf("please register at the leader before requesting votes")
	}

	voteGranted := true

	state.DefaultPersistentState.Mutex.Lock()
	defer state.DefaultPersistentState.Mutex.Unlock()

	if ( // check if already voted this term
		state.DefaultPersistentState.VoteFor.Term == v.Term && state.DefaultPersistentState.VoteFor.Id != nil) ||
		// check if i'm in a newer term already
		state.DefaultPersistentState.CurrentTerm > v.Term ||
		// check if this node has newer log entries
		state.DefaultPersistentState.GetLastLogIndexFragile() > v.LastLogIndex ||
		// check if the requesting node has an outdated term
		state.DefaultPersistentState.GetLastLogTermFragile() > v.LastLogTerm {
		voteGranted = false
		lg.Log.Infof("Denied vote request from %d in term %d with LastLogIndex %d and LastLogTerm %d", v.CandidateId, v.Term, v.LastLogIndex, v.LastLogTerm)

		if state.DefaultPersistentState.CurrentSate == state.Leader {
			// got message from client, reset connection backoff
			lg.Log.Debug("Reset connection backoff")
			client.GetClientSet().ResetBackoff()
		}

		if state.DefaultPersistentState.CurrentTerm < v.Term {
			state.DefaultPersistentState.CurrentTerm = v.Term
			state.DefaultPersistentState.CurrentSate = state.Follower
		}
	} else {
		// elect and update self
		voteGranted = true
		candidateId := v.CandidateId
		state.DefaultPersistentState.UpdateFragile(state.Follower, v.Term, &candidateId)
		lg.Log.Infof("Voted for %d in term %d with LastLogIndex %d and LastLogTerm %d", v.CandidateId, v.Term, v.LastLogIndex, v.LastLogTerm)
	}

	return &rpc.VoteResponse{
		Term:        state.DefaultPersistentState.CurrentTerm,
		VoteGranted: voteGranted,
	}, nil
}
