package state

import (
	"fmt"
	"github.com/LaCodon/verteilte-systeme/internal/helper"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
	firstNewElementIndex := 0

	// remove all inconsistent elements
	for i, element := range elements {
		if len(s.Log) > int(element.Index) {
			if element.Term != s.Log[element.Index].Term {
				firstNewElementIndex = i
				s.Log = s.Log[:element.Index]
				break
			} else {
				firstNewElementIndex = i+1
			}
		}
	}

	// add all new elements
	if firstNewElementIndex < len(elements) {
		lg.Log.Debugf("Adding following entries to log: %v", elements[firstNewElementIndex:])
		s.Log = append(s.Log, elements[firstNewElementIndex:]...)
	}
}
func (s *PersistentState) ApplyLogToStateMachine(lastApplied int32, commitIndex int32) int32 {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	return s.ApplyLogToStateMachineFragile(lastApplied, commitIndex)
}

func (s *PersistentState) ApplyLogToStateMachineFragile(lastApplied int32, commitIndex int32) int32 {
	file, err := os.OpenFile(helper.GetLogFilePath(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		lg.Log.Warningf("Error while opening file: %s", err)
	}
	defer file.Close()

	for i:=int32(0); i<=commitIndex; i++ {
		if s.Log[i].Index <= lastApplied{
			continue
		}
		lg.Log.Debugf("Writing log entry to file: " + config.Default.LogFormatString, s.Log[i].Index, s.Log[i].Term, s.Log[i].Action, s.Log[i].Key, s.Log[i].Value)
		_, err := file.WriteString(fmt.Sprintf(config.Default.LogFormatString, s.Log[i].Index, s.Log[i].Term, s.Log[i].Action, s.Log[i].Key, s.Log[i].Value))
		if err != nil {
			lg.Log.Error(err)
			break
		}
		lastApplied++
	}
	return lastApplied
}

func (s *PersistentState) InitLog() {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	//read data from file
	data, ioErr := ioutil.ReadFile(helper.GetLogFilePath())
	if ioErr != nil {
		lg.Log.Warningf("Error while reading file: %s", ioErr)
	}

	s.Log = []*rpc.LogEntry{}
	//parse entries
	entries := strings.Split(string(data), "\n")
	for _, entry := range entries {
		lg.Log.Debugf("Loading log entry: %s",  entry)
		parts := strings.Fields(entry)

		if len(parts) != 5 {
			// format is not correct, ignoring line
			lg.Log.Warningf("log file contains incorrect formatted line: %s", entry)
			continue
		}

		//parse integer values
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			lg.Log.Fatalf("could not initialize log! \"%s\" is no integer", parts[0])
		}
		term, err := strconv.Atoi(parts[1])
		if err != nil {
			lg.Log.Fatalf("could not initialize log! \"%s\" is no integer", parts[1])
		}
		action, err := strconv.Atoi(parts[2])
		if err != nil {
			lg.Log.Fatalf("could not initialize log! \"%s\" is no integer", parts[2])
		}

		//append entry to log
		logEntry := &rpc.LogEntry{
			Index:  int32(index),
			Term:   int32(term),
			Action: int32(action),
			Key:    parts[3],
			Value:  parts[4],
		}
		lg.Log.Debugf("Log entry %d loaded: %v", logEntry.Index, logEntry)
		s.Log = append(s.Log, logEntry)
	}

	if len(s.Log) > 0 {
		s.CurrentTerm = s.GetLastLogTermFragile()
		DefaultVolatileState.LastApplied = s.Log[len(s.Log)-1].Index
	}
}

func (s *PersistentState) GetCurrentStateMachineFragile() map[string]string {
	var m map[string]string
	for _, e := range s.Log {
		if e.Action == 1 {
			//SET
			m[e.Key] = e.Value
		} else if e.Action == 2 {
			//DELETE
			delete(m, e.Key)
		}
	}
	return m
}