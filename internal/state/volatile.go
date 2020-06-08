package state

type VolatileState struct {
	State
	commitIndex int32
	lastApplied int32
}
