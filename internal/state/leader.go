package state

type LeaderState struct {
	State
	NextIndex  map[int]int
	MatchIndex map[int]int
}

var DefaultLeaderState *LeaderState