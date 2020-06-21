package state

import "github.com/LaCodon/verteilte-systeme/pkg/redolog"

func init() {
	DefaultPersistentState = &PersistentState{
		CurrentSate: Follower,
		CurrentTerm: 1,
		VoteFor: struct {
			Id   *int32
			Term int32
		}{Id: nil, Term: 1},
		Log: []*redolog.Element{},
	}
}
