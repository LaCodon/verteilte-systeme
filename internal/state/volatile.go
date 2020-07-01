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

func (s *VolatileState) SetCommitIndex(i int32) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.CommitIndex = i
}

func (s *VolatileState) GetCommitIndex() int32 {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	return s.CommitIndex
}
