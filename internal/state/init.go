package state

import "github.com/LaCodon/verteilte-systeme/pkg/rpc"

func init() {
	DefaultPersistentState = &PersistentState{
		CurrentSate: Follower,
		CurrentTerm: 1,
		VoteFor: struct {
			Id   *uint32
			Term int32
		}{Id: nil, Term: 1},
		Log: []*rpc.LogEntry{},
	}

	DefaultLeaderState = &LeaderState{
		NextIndex:  make(map[uint32]int32),
		MatchIndex: make(map[uint32]int32),
	}

	DefaultVolatileState = &VolatileState{
		CommitIndex:   -1,
		LastApplied:   -1,
		CurrentLeader: "",
	}
}
