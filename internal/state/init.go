package state

import "github.com/LaCodon/verteilte-systeme/pkg/rpc"

func init() {
	DefaultPersistentState = &PersistentState{
		CurrentSate: Follower,
		CurrentTerm: 1,
		VoteFor: struct {
			Id   *int32
			Term int32
		}{Id: nil, Term: 1},
		Log: []*rpc.LogEntry{},
	}

	DefaultLeaderState = &LeaderState{
		NextIndex:  make(map[int]int),
		MatchIndex: make(map[int]int),
	}

	DefaultVolatileState = &VolatileState{
		CommitIndex:   0,
		LastApplied:   0,
		CurrentLeader: "",
	}
}
