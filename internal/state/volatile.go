package state

type VolatileState struct {
	State
	CommitIndex int32
	LastApplied int32
}

var DefaultVolatileState *VolatileState



// Makes no use of RW-Mutex.
func (s *VolatileState) IncreaseLastAppliedFragile() {
	s.LastApplied++
}