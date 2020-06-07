package state

import "github.com/LaCodon/verteilte-systeme/pkg/logentry"

type PersistentState struct {
	*State
	currentTerm int32
	voteFor     int32
	log         []*logentry.Element
}

func (s *PersistentState) SetCurrentTerm(t int32) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.currentTerm = t
}

func (s *PersistentState) SetVoteFor(v int32) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.voteFor = v
}

func (s *PersistentState) AddToLog(l *logentry.Element) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.log = append(s.log, l)
}

func (s *PersistentState) GetCurrentTerm() int32 {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	return s.currentTerm
}

func (s *PersistentState) GetVoteFor() int32 {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	return s.voteFor
}

func (s *PersistentState) GetLogElement(i int) logentry.Element {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	// create copy to prevent manipulation
	r := *s.log[i]
	return r
}

func (s *PersistentState) GetLogLength() int {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	return len(s.log)
}
