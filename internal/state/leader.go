package state

type LeaderState struct {
	*State
	nextIndex  map[int]int
	matchIndex map[int]int
}
