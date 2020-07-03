package state

type LeaderState struct {
	State
	NextIndex  map[uint32]int32
	MatchIndex map[uint32]int32
}

var DefaultLeaderState *LeaderState
