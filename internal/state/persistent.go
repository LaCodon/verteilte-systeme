package state

import (
	"github.com/LaCodon/verteilte-systeme/pkg/redolog"
)

type NodeState int

const (
	Follower  NodeState = iota + 1
	Candidate NodeState = iota + 1
	Leader    NodeState = iota + 1
)

var DefaultPersistentState *PersistentState

type PersistentState struct {
	State
	CurrentSate NodeState
	CurrentTerm int32
	VoteFor     struct {
		Id   *int32
		Term int32
	}
	Log []*redolog.Element
}

// Makes no use of RW-Mutex.
func (s *PersistentState) GetLastLogIndexFragile() int32 {
	return int32(len(s.Log) - 1)
}

// UpdateFragile updates all relevant persistent state variables.
// Makes no use of RW-Mutex.
func (s *PersistentState) UpdateFragile(newState NodeState, newTerm int32, newVoteFor *int32) {
	s.CurrentSate = newState
	s.CurrentTerm = newTerm
	s.VoteFor.Term = newTerm
	s.VoteFor.Id = newVoteFor
}

// Makes no use of RW-Mutex.
func (s *PersistentState) GetLastLogTermFragile() int32 {
	length := len(s.Log)
	lastIndex := length - 1

	if length == 0 {
		return 0
	}

	return s.Log[lastIndex].Term
}

func (s *PersistentState) SetCurrentState(n NodeState) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.CurrentSate = n
}

func (s *PersistentState) IncrementCurrentTerm() {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.CurrentTerm++
}

func (s *PersistentState) SetCurrentTerm(t int32) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.CurrentTerm = t
}

func (s *PersistentState) AddToLog(l *redolog.Element) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Log = append(s.Log, l)
}

func (s *PersistentState) GetCurrentState() NodeState {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	return s.CurrentSate
}

func (s *PersistentState) GetCurrentTerm() int32 {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	return s.CurrentTerm
}

func (s *PersistentState) GetLogElement(i int) redolog.Element {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	// create copy to prevent manipulation
	r := *s.Log[i]
	return r
}

func (s *PersistentState) GetLogLength() int {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	return len(s.Log)
}
