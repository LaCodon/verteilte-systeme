package client

import (
	"context"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/redolog"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"math/rand"
	"strconv"
	"time"
)

func BeLeader(ctx context.Context) {

	SendHeartbeatsAsync(ctx)

	HandleUserInput(ctx)

}

func SendHeartbeatsAsync(ctx context.Context) {
	go func(ctx context.Context) {
		clients := ConnectToNodes(config.Default.PeerNodes.Value())
		for state.DefaultPersistentState.GetCurrentState() == state.Leader && ctx.Err() == nil {
			term := state.DefaultPersistentState.GetCurrentTerm()

			for _, client := range clients {
				go func(c rpc.NodeClient, ctx context.Context) {
					ctx, _ = context.WithTimeout(ctx, 500*time.Millisecond)
					_, err := c.AppendEntries(ctx, &rpc.AppendEntriesRequest{Term: term})
					if err != nil {
						lg.Log.Debugf("error from client.AppendEntries: %s", err)
					}
				}(client, ctx)
			}

			time.Sleep(800 * time.Millisecond)
		}
	}(ctx)
}

func HandleUserInput(ctx context.Context) {
	lastLogIndex := state.DefaultPersistentState.GetLastLogIndexFragile()
	lastLogTerm := state.DefaultPersistentState.GetLastLogTermFragile()

	//TODO get user Input
	keys := []string{"a","b","c","x","y","z"}
	randomKey := rand.Intn(len(keys))
	randomValue := rand.Intn(50)

	var randomAction redolog.Action
	if(rand.Intn(2) == 0) {
		randomAction = redolog.ActionSet
	} else {
		randomAction = redolog.ActionDelete
	}

	// handle user Input
	element := &redolog.Element{
		Index: 				  state.DefaultPersistentState.GetLastLogIndexFragile()+1,
		Term:                 state.DefaultPersistentState.CurrentTerm,
		Key:                  keys[randomKey],
		Action:               randomAction,
		Value:                strconv.Itoa(randomValue),
	}

	//append user input to log
	state.DefaultPersistentState.AddToLog(element)
	state.DefaultVolatileState.IncreaseLastAppliedFragile()

	var entries []*rpc.LogEntry
	entries = append(entries, redolog.ElementToLogEntry(element))

	//TODO send user Input to all follower
	r := &rpc.AppendEntriesRequest{
		Term:                 state.DefaultPersistentState.CurrentTerm,
		LeaderId:             int32(config.Default.NodeId),
		PrevLogIndex:         lastLogIndex,
		PrevLogTerm:          lastLogTerm,
		Entries:              entries,
		LeaderCommit:         state.DefaultVolatileState.CommitIndex(),
	}
}
