package server

import (
	"context"
	"fmt"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/client"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/redolog"
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

	// check lastLogIndex and Term
	if !state.DefaultPersistentState.ContainsLogElementFragile(ar.PrevLogIndex, ar.PrevLogTerm) {
		return &rpc.AppendEntriesResponse{
			Term: state.DefaultPersistentState.GetCurrentTerm(),
			Success: false,
		}, fmt.Errorf("previous log entry not found")
	}

	// update log
	var elements []*redolog.Element
	for _, e := range ar.Entries{
		elements = append(elements, redolog.LogEntryToElement(e))
	}
	state.DefaultPersistentState.UpdateAndAppendLog(elements)

	if ar.LeaderCommit > state.DefaultVolatileState.CommitIndex {
		if ar.LeaderCommit < elements[len(elements)-1].Index {
			state.DefaultVolatileState.CommitIndex = ar.LeaderCommit
		} else {
			state.DefaultVolatileState.CommitIndex = elements[len(elements)-1].Index
		}
	}

	return &rpc.AppendEntriesResponse{
		Term: state.DefaultPersistentState.GetCurrentTerm(),
	}, nil
}
