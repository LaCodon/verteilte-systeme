package state

import "sync"

type State struct {
	Mutex sync.RWMutex
}
