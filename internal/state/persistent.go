package state

import (
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
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
		Id   *uint32
		Term int32
	}
	Log []*rpc.LogEntry
}

// Makes no use of RW-Mutex.
func (s *PersistentState) GetLastLogIndexFragile() int32 {
	return int32(len(s.Log) - 1)
}

// UpdateFragile updates all relevant persistent state variables.
// Makes no use of RW-Mutex.
func (s *PersistentState) UpdateFragile(newState NodeState, newTerm int32, newVoteFor *uint32) {
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

func (s *PersistentState) ContainsLogElementFragile(index int32, term int32) bool {
	if index == -1 {
		return true
	}
	if int(index) >= len(s.Log) {
		return false
	}
	return s.Log[index].Term == term
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

func (s *PersistentState) AddToLogFragile(l ...*rpc.LogEntry) {
	for _, e := range l {
		s.Log = append(s.Log, e)
	}
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

func (s *PersistentState) GetLogElement(i int) rpc.LogEntry {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	// create copy to prevent manipulation
	r := rpc.LogEntry{
		Index:  s.Log[i].Index,
		Term:   s.Log[i].Term,
		Key:    s.Log[i].Key,
		Action: s.Log[i].Action,
		Value:  s.Log[i].Value,
	}
	return r
}

func (s *PersistentState) GetLogLength() int {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	return len(s.Log)
}

// UpdateAndAppendLogFragile removes all invalid elements and appends the new received ones
func (s *PersistentState) UpdateAndAppendLogFragile(elements []*rpc.LogEntry) {
	//commitIndex := DefaultVolatileState.GetCommitIndex()

	//for _, element := range elements {
	//	if element.Index > commitIndex {
	//		s.Log = append(s.Log, element)
	//		if element.Index != int32(len(s.Log)-1) {
	//			lg.Log.Errorf("Appended bullshit data!")
	//		}
	//	}
	//}

	var firstNewElementIndex int32

	// remove all inconsistent elements
	for _, element := range elements {
		if len(s.Log) > int(element.Index) {
			if element.Term != s.Log[element.Index].Term {
				firstNewElementIndex = element.Index
				s.Log = s.Log[:element.Index]
				break
			}
		}
	}

	// add all new elements
	s.Log = append(s.Log, elements[firstNewElementIndex:]...)
}
