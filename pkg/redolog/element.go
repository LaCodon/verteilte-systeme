package redolog

type Action int

const (
	ActionSet    Action = iota + 1
	ActionDelete Action = iota + 1
)

type Element struct {
	Term   int32
	Key    string
	Action Action
	Value  string
}
