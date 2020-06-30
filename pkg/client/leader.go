package client

import (
	"context"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var NewLogEntries chan *rpc.LogEntry

func BeLeader(ctx context.Context) {
	clients := ConnectToNodes(config.Default.PeerNodes.Value())

	// reset volatile leader state
	for i := range clients {
		state.DefaultLeaderState.NextIndex[i] = int(state.DefaultPersistentState.GetLastLogIndexFragile()) + 1
		state.DefaultLeaderState.MatchIndex[i] = 0
	}

	SendHeartbeatsAsync(ctx, clients)

	HandleUserInput(ctx, clients)

}

func SendHeartbeatsAsync(ctx context.Context, clients ClientSet) {
	go func(ctx context.Context, clients ClientSet) {
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
	}(ctx, clients)
}

func Send(ctx context.Context, clients ClientSet) {
	go func(ctx context.Context, clients ClientSet) {
		for state.DefaultPersistentState.GetCurrentState() == state.Leader && ctx.Err() == nil {

			//handle new User input
			var entries []*rpc.LogEntry
			newUserInput := true
			for newUserInput {
				select {
				case newEntry := <-NewLogEntries:
					entries = append(entries, newEntry)
				default:
					newUserInput = false
				}
			}

			// generate request
			state.DefaultPersistentState.Mutex.RLock()
			state.DefaultVolatileState.Mutex.RLock()
			r := &rpc.AppendEntriesRequest{
				Term:                 state.DefaultPersistentState.CurrentTerm,
				LeaderId:             int32(config.Default.NodeId),
				PrevLogIndex:         state.DefaultPersistentState.GetLastLogIndexFragile(),
				PrevLogTerm:          state.DefaultPersistentState.GetLastLogTermFragile(),
				Entries:              entries,
				LeaderCommit:         state.DefaultVolatileState.CommitIndex,
			}
			state.DefaultPersistentState.AddToLog(entries...)
			state.DefaultVolatileState.Mutex.RUnlock()
			state.DefaultPersistentState.Mutex.RUnlock()

			// count the successful replicated logs; start at 1 for leaders log
			successfulReplications := 1
			nodeCount := len(config.Default.PeerNodes.Value()) + 1
			var lock sync.Mutex

			//sending request to all nodes
			for i, client := range clients {
				id := i
				req := r
				logReplicated := false
				go func(c rpc.NodeClient, ctx context.Context) {
					for !logReplicated {
						ctx, _ = context.WithTimeout(ctx, 500*time.Millisecond)
						resp, err := c.AppendEntries(ctx, r)
						if err == nil {
							if resp.Success {
								logReplicated = true
								lock.Lock()
								successfulReplications++
								lock.Unlock()

								if successfulReplications > nodeCount/2 {
									state.DefaultVolatileState.CommitIndex = state.DefaultPersistentState.GetLastLogIndexFragile()
									lg.Log.Debugf("Replicated log on the majority of the nodes, new commit index: %s", state.DefaultVolatileState.CommitIndex)
								}
							} else {
								if state.DefaultLeaderState.NextIndex[id]>1 {
									state.DefaultLeaderState.NextIndex[id]--
									req.Entries = state.DefaultPersistentState.Log[state.DefaultLeaderState.NextIndex[id]:]
								}
								lg.Log.Debugf("Replicating log failed for node %d, retrying with NextIndex=%d", id, state.DefaultLeaderState.NextIndex[id])
							}
						} else {
							lg.Log.Debugf("error from client.AppendEntries: %s", err)
						}
					}
				}(client, ctx)
			}
			//TODO: assure that ALL nodes have the correct log

			time.Sleep(800 * time.Millisecond)
		}
	}(ctx, clients)
}

func HandleUserInput(ctx context.Context, clients ClientSet) {
	lastLogIndex := state.DefaultPersistentState.GetLastLogIndexFragile()
	lastLogTerm := state.DefaultPersistentState.GetLastLogTermFragile()

	//TODO get user Input
	keys := []string{"a","b","c","x","y","z"}
	randomKey := rand.Intn(len(keys))
	randomValue := rand.Intn(50)

	randomAction := rand.Int31n(2)

	// handle user Input
	entry := &rpc.LogEntry{
		Index: 				  state.DefaultPersistentState.GetLastLogIndexFragile()+1,
		Term:                 state.DefaultPersistentState.CurrentTerm,
		Key:                  keys[randomKey],
		Action:               randomAction,
		Value:                strconv.Itoa(randomValue),
	}

	//append user input to log
	state.DefaultPersistentState.AddToLog(entry)
	state.DefaultVolatileState.IncreaseLastAppliedFragile()
	var entries []*rpc.LogEntry
	entries = append(entries, entry)

	//TODO send user Input to all follower
	nodeCount := len(config.Default.PeerNodes.Value()) + 1
	incomingResponses := make(chan *rpc.AppendEntriesResponse, nodeCount)

	state.DefaultPersistentState.Mutex.RLock()
	r := &rpc.AppendEntriesRequest{
		Term:                 state.DefaultPersistentState.CurrentTerm,
		LeaderId:             int32(config.Default.NodeId),
		PrevLogIndex:         lastLogIndex,
		PrevLogTerm:          lastLogTerm,
		Entries:              entries,
		LeaderCommit:         state.DefaultVolatileState.CommitIndex,
	}
	state.DefaultPersistentState.Mutex.RUnlock()

	lg.Log.Debugf("Sending UserInput for term %d, LeaderID %d, PrevLogIndex %d, PrevLogTerm %d, LeaderCommit %d", r.Term, r.LeaderId, r.PrevLogIndex, r.PrevLogTerm, r.LeaderCommit)
	lg.Log.Debugf("Log Element: %v", entries)


	// send log to all other nodes
	for _, c := range clients {
		go func(client rpc.NodeClient) {
			ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)
			resp, err := client.AppendEntries(ctx, r)
			if err == nil {
				if resp.Success {
					incomingResponses <- resp
				}
			} else {
				lg.Log.Debugf("Got error from client.AppendEntries call: %s", err)
			}
		}(c)
	}
	
	replicating := true
	//replicatedLogs start at 1 because leader has it already
	replicatedLogs := 1

	for replicating {
		select {
		case <-timedOut:
			lg.Log.Info("timed out waiting for AppendEntries responds")
			replicating = false
		case resp := <-incomingResponses:
			lg.Log.Infof("Got AppendEntries response: %s", resp)
			if resp.Success {
				replicatedLogs++

				if replicatedLogs > nodeCount/2 {
					state.DefaultVolatileState.CommitIndex = state.DefaultPersistentState.GetLastLogIndexFragile()
				}
			}
		}
	}
}
