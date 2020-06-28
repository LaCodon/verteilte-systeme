package redolog

import (
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
)

type Action int

const (
	ActionSet    Action = iota + 1
	ActionDelete Action = iota + 1
)

type Element struct {
	Index  int32
	Term   int32
	Key    string
	Action Action
	Value  string
}

func ElementToLogEntry(element *Element) *rpc.LogEntry {
	return &rpc.LogEntry{
		Index:                element.Index,
		Term:                 element.Term,
		Key:                  element.Key,
		Action:               int32(element.Action),
		Value:                element.Value,
	}
}

func LogEntryToElement(entry *rpc.LogEntry) *Element {
	var action Action
	if entry.Action == 1 {
		action = ActionSet
	} else {
		action = ActionDelete
	}

	return &Element{
		Index:  entry.Index,
		Term:   entry.Term,
		Key:    entry.Key,
		Action: action,
		Value:  entry.Value,
	}
}
