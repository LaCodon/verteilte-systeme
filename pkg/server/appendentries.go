package server

import (
	"context"
	"fmt"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/client"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
)

// AppendEntries gets called by the leader for heartbeats / new data
func (s *Server) AppendEntries(c context.Context, ar *rpc.AppendEntriesRequest) (*rpc.AppendEntriesResponse, error) {
	lg.Log.Debugf("Got heartbeat with term %d", ar.Term)

	if state.DefaultPersistentState.GetCurrentTerm() > ar.Term {
		return &rpc.AppendEntriesResponse{
			Term: state.DefaultPersistentState.GetCurrentTerm(),
			Success: false,
		}, fmt.Errorf("old term, ignored request")
	}

	// immediately convert to follower (and mark leader as voted to prevent from new leader in this term)
	state.DefaultPersistentState.UpdateFragile(state.Follower, ar.Term, &ar.LeaderId)

	// send heartbeat to other go routines
	client.Heartbeat <- true

	// check PrevLogIndex and Term
	if !state.DefaultPersistentState.ContainsLogElementFragile(ar.PrevLogIndex, ar.PrevLogTerm) {
		return &rpc.AppendEntriesResponse{
			Term: state.DefaultPersistentState.GetCurrentTerm(),
			Success: false,
		}, fmt.Errorf("previous log entry not found")
	}

	// update log
	if len(ar.Entries) > 0 {
		lg.Log.Infof("received log entries: %v", ar.Entries)
		state.DefaultPersistentState.UpdateAndAppendLog(ar.Entries)

		lg.Log.Debug("setting commit index...")
		if ar.LeaderCommit > state.DefaultVolatileState.CommitIndex {
			if ar.LeaderCommit < ar.Entries[len(ar.Entries)-1].Index {
				state.DefaultVolatileState.CommitIndex = ar.LeaderCommit
			} else {
				state.DefaultVolatileState.CommitIndex = ar.Entries[len(ar.Entries)-1].Index
			}
		}
	}

	return &rpc.AppendEntriesResponse{
		Term: state.DefaultPersistentState.GetCurrentTerm(),
		Success:true,
	}, nil
}
