package server

import (
	"context"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/client"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
)

// AppendEntries gets called by the leader for heartbeats / new data
func (s *Server) AppendEntries(c context.Context, ar *rpc.AppendEntriesRequest) (*rpc.AppendEntriesResponse, error) {
	lg.Log.Debugf("Got heartbeat with term %d", ar.Term)

	if state.DefaultPersistentState.GetCurrentTerm() > ar.Term {
		lg.Log.Debugf("Ignored heartbeat because it had outdated term")
		return &rpc.AppendEntriesResponse{
			Term:    state.DefaultPersistentState.GetCurrentTerm(),
			Success: false,
		}, nil
	}

	state.DefaultPersistentState.Mutex.Lock()
	defer state.DefaultPersistentState.Mutex.Unlock()

	// immediately convert to follower (and mark leader as voted to prevent from new leader in this term)
	state.DefaultPersistentState.UpdateFragile(state.Follower, ar.Term, &ar.LeaderId)

	// send heartbeat to other go routines
	client.Heartbeat <- true

	if config.Default.HasNodesDiff(ar.AllNodes) {
		// save new peer node information
		config.Default.SetNewNodes(ar.AllNodes)
		client.ForceClientReconnect = true
	}

	lg.Log.Debugf("PrevLogIndex: %d ; PrevLogTerm: %d", ar.PrevLogIndex, ar.PrevLogTerm)

	// check PrevLogIndex and Term
	if !state.DefaultPersistentState.ContainsLogElementFragile(ar.PrevLogIndex, ar.PrevLogTerm) {
		lg.Log.Debugf("Ignored AppendEntries for index %d because i'm missing older entries", ar.LeaderCommit)
		return &rpc.AppendEntriesResponse{
			Term:    state.DefaultPersistentState.CurrentTerm,
			Success: false,
		}, nil
	}

	lg.Log.Debugf("old commit index: %d", state.DefaultVolatileState.GetCommitIndex())
	// update log
	if len(ar.Entries) > 0 {
		lg.Log.Infof("received log entries: %v", ar.Entries)
		state.DefaultPersistentState.UpdateAndAppendLogFragile(ar.Entries)

		if ar.LeaderCommit < ar.Entries[len(ar.Entries)-1].Index {
			// sync with leader commit
			state.DefaultVolatileState.SetCommitIndex(ar.LeaderCommit)
		} else {
			// set commit index to highest element I received
			state.DefaultVolatileState.SetCommitIndex(ar.Entries[len(ar.Entries)-1].Index)
		}
	} else {
		state.DefaultVolatileState.SetCommitIndex(ar.LeaderCommit)
	}
	lg.Log.Debugf("new commit index: %d", state.DefaultVolatileState.GetCommitIndex())

	state.DefaultVolatileState.SetCurrentLeader(ar.LeaderTarget)

	return &rpc.AppendEntriesResponse{
		Term:    state.DefaultPersistentState.CurrentTerm,
		Success: true,
	}, nil
}
