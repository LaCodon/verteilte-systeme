package state

import "github.com/LaCodon/verteilte-systeme/pkg/redolog"

func init() {
	DefaultPersistentState = &PersistentState{
		CurrentSate: Follower,
		CurrentTerm: 1,
		VoteFor:     nil,
		Log:         []*redolog.Element{},
	}
}
