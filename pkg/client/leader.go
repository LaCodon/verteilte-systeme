package client

import (
	"context"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/lg"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type UserInput struct{
	Key string
	Var string
	Action int32
}

var NewUserInput chan *UserInput

func BeLeader(ctx context.Context) {
	clients := DefaultClientSet
	// reset volatile leader state
	for i := range clients {
		state.DefaultLeaderState.NextIndex[i] = int(state.DefaultPersistentState.GetLastLogIndexFragile()) + 1
		state.DefaultLeaderState.MatchIndex[i] = 0
		lg.Log.Debugf("initialised nextIndex[%d]=%d", i, state.DefaultLeaderState.NextIndex[i])
	}



	//read user Input
	NewUserInput = make(chan *UserInput, 20)
	go HandleUserInput()

	Send(ctx, clients)


}

func HandleUserInput() {
	for {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			lg.Log.Warning(err)
		}
		file := path.Join(dir, "../userInput/input.txt")
		//read data from file
		data, err := ioutil.ReadFile(file)
		if err != nil {
			lg.Log.Warningf("Error while reading file: %s", err)
		}
		// remove data from file
		inputFile, err := os.OpenFile(file, os.O_TRUNC, 0666)
		if err != nil {
			lg.Log.Warningf("Error while reading file: %s", err)
		}
		inputFile.Close()
		//parse entries
		entries := strings.Split(string(data), "\n")
		for _, entry := range entries {
			cmd := strings.Fields(entry)
			if len(cmd) > 1 {
				if len(cmd) == 2 {
					cmd = append(cmd, "0")
				}
				action, err := strconv.Atoi(cmd[0])
				lg.Log.Infof("Recieved user command: %d %s %s", action, cmd[1], cmd[2])
				if err == nil {
					NewUserInput <- &UserInput{
						Key:    cmd[1],
						Var:    cmd[2],
						Action: int32(action),
					}
				} else {
					lg.Log.Warningf("Could not convert user input \"%s\" to action: %s", cmd[0], err)
				}
			}
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func Send(ctx context.Context, clients ClientSet) {
	for state.DefaultPersistentState.GetCurrentState() == state.Leader && ctx.Err() == nil {
		go func(ctx context.Context, clients ClientSet) {
			lg.Log.Debug("Sending next AppendEntries request.")
			//handle new User input
			var entries []*rpc.LogEntry
			moreUserInput := true
			for moreUserInput {
				select {
				case input := <- NewUserInput:
					newEntry := &rpc.LogEntry{
						Index:  state.DefaultPersistentState.GetLastLogIndexFragile() + 1,
						Term:   state.DefaultPersistentState.CurrentTerm,
						Key:    input.Key,
						Action: input.Action,
						Value:  input.Var,
					}
					lg.Log.Debugf("new user entry: %s", newEntry)
					entries = append(entries, newEntry)
				default:
					lg.Log.Debug("no new user entries")
					moreUserInput = false
				}
			}
			lg.Log.Debugf("New Entries: %v", entries)
			LastLogIndex := state.DefaultPersistentState.GetLastLogIndexFragile()
			LastLogTerm := state.DefaultPersistentState.GetLastLogTermFragile()
			state.DefaultPersistentState.AddToLog(entries...)

			// count the successful replicated logs; start at 1 for leaders log
			successfulReplications := 1
			nodeCount := len(config.Default.PeerNodes.Value()) + 1
			var lock sync.Mutex

			//sending request to all nodes
			for i, client := range clients {
				id := i
				entriesToSend := state.DefaultPersistentState.Log[state.DefaultLeaderState.NextIndex[id]:]
				r := &rpc.AppendEntriesRequest{
					Term:         state.DefaultPersistentState.CurrentTerm,
					LeaderId:     int32(config.Default.NodeId),
					PrevLogIndex: LastLogIndex,
					PrevLogTerm:  LastLogTerm,
					Entries:      entriesToSend,
					LeaderCommit: state.DefaultVolatileState.CommitIndex,
				}
				go func(c rpc.NodeClient, ctx context.Context) {
					if len(entriesToSend) == 0 {
						lg.Log.Infof("sending heartbeat to node %d", id)
					} else {
						lg.Log.Infof("sending entries to node %d: %v", id, entriesToSend)
					}
					ctx, _ = context.WithTimeout(ctx, 500*time.Millisecond)
					resp, err := c.AppendEntries(ctx, r)
					if err == nil {
						lg.Log.Infof("node %d responded: %s", id, resp)
						if resp.Success {
							lock.Lock()
							successfulReplications++
							lock.Unlock()
							state.DefaultLeaderState.NextIndex[id] = int(state.DefaultPersistentState.GetLastLogIndexFragile()) + 1

							if successfulReplications > nodeCount/2 {
								state.DefaultVolatileState.CommitIndex = state.DefaultPersistentState.GetLastLogIndexFragile()
								lg.Log.Debugf("Replicated log on the majority of the nodes, new commit index: %d", state.DefaultVolatileState.CommitIndex)
							}
						} else {
							if resp.Term > state.DefaultPersistentState.CurrentTerm {
								state.DefaultPersistentState.SetCurrentState(state.Follower)
							} else {
								if state.DefaultLeaderState.NextIndex[id] > 1 {
									state.DefaultLeaderState.NextIndex[id]--
								}
								lg.Log.Debugf("Replicating log failed for node %d, retrying with NextIndex=%d", id, state.DefaultLeaderState.NextIndex[id])
							}
						}
					} else {
						lg.Log.Debugf("error from client.AppendEntries: %s", err)
					}
				}(client, ctx)
			}

		}(ctx, clients)
		time.Sleep(800 * time.Millisecond)
	}
}
