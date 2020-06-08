/*
The state package is used as a simple store for the current states of the node. It provides a RWMutex per
state to enable concurrent use of it

State types

The raft consensus algorithm uses three different state types for various situations in a nodes lifespan:
Persistent State, Volatile State and Leader State.

*/
package state
