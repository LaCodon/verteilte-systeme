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

type UserInput struct {
	Key    string
	Var    string
	Action int32
}

var NewUserInput chan *UserInput

func BeLeader(ctx context.Context) {
	// reset volatile leader state
	for i := range GetClientSet() {
		state.DefaultLeaderState.NextIndex[i] = int(state.DefaultPersistentState.GetLastLogIndexFragile()) + 1
		state.DefaultLeaderState.MatchIndex[i] = 0
		lg.Log.Debugf("initialised nextIndex[%d]=%d", i, state.DefaultLeaderState.NextIndex[i])
	}

	//read user Input
	NewUserInput = make(chan *UserInput, 20)
	go HandleUserInput()

	Send(ctx)
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
					lg.Log.Debugf("-----------------send input---------------------------")
				} else {
					lg.Log.Warningf("Could not convert user input \"%s\" to action: %s", cmd[0], err)
				}
			}
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func Send(ctx context.Context) {
	for state.DefaultPersistentState.GetCurrentState() == state.Leader && ctx.Err() == nil {
		lg.Log.Debug("Sending next AppendEntries request.")
		//handle new User input
		var entries []*rpc.LogEntry
		moreUserInput := true
		for moreUserInput {
			select {
			case input := <-NewUserInput:
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

		var (
			PrevLogIndex    int32
			PrevLogTerm     int32
			CurrentLogIndex int32
			CurrentTerm     int32
		)
		{
			state.DefaultPersistentState.Mutex.Lock()
			PrevLogIndex = state.DefaultPersistentState.GetLastLogIndexFragile()
			PrevLogTerm = state.DefaultPersistentState.GetLastLogTermFragile()
			state.DefaultPersistentState.AddToLogFragile(entries...)
			CurrentLogIndex = state.DefaultPersistentState.GetLastLogIndexFragile()
			CurrentTerm = state.DefaultPersistentState.CurrentTerm

			state.DefaultPersistentState.Mutex.Unlock()
		}

		// count the successful replicated logs; start at 1 for leaders log
		successfulReplications := 1
		nodeCount := config.Default.PeerNodeCount() + 1
		var lock sync.Mutex

		//sending request to all nodes
		for id, client := range GetClientSet() {
			state.DefaultPersistentState.Mutex.RLock()
			entriesToSend := state.DefaultPersistentState.Log[state.DefaultLeaderState.NextIndex[id]:]
			state.DefaultPersistentState.Mutex.RUnlock()

			r := &rpc.AppendEntriesRequest{
				Term:         CurrentTerm,
				LeaderId:     int32(config.Default.NodeId),
				PrevLogIndex: PrevLogIndex,
				PrevLogTerm:  PrevLogTerm,
				Entries:      entriesToSend,
				LeaderCommit: state.DefaultVolatileState.GetCommitIndex(),
				AllNodes:     config.Default.AllNodes.Value(),
				LeaderTarget: config.Default.MyNode,
			}
			go func(c *Client, parentCtx context.Context, clientIndex int) {
				if len(entriesToSend) == 0 {
					lg.Log.Debugf("sending heartbeat to node %d", clientIndex)
				} else {
					lg.Log.Debugf("sending entries to node %d: %v", clientIndex, entriesToSend)
				}
				appendCtx, _ := context.WithTimeout(parentCtx, config.Default.AppendEntriesTimeout)
				resp, err := c.NodeClient.AppendEntries(appendCtx, r)
				if err != nil {
					lg.Log.Debugf("error from Client.AppendEntries: %s", err)
					c.ErrorCount++

					if c.ErrorCount > config.Default.KickThreshold {
						config.Default.RemoveNode(c.Target)
						ForceClientReconnect = true
						lg.Log.Warningf("removed node %s from cluster", c.Target)
					}

					return
				}

				// got response, reset error counter
				c.ErrorCount = 0

				lg.Log.Debugf("node %d responded: %s", clientIndex, resp)
				if resp.Success {
					lock.Lock()
					successfulReplications++
					lock.Unlock()

					state.DefaultLeaderState.Mutex.Lock()
					state.DefaultLeaderState.NextIndex[clientIndex] = int(CurrentLogIndex) + 1
					state.DefaultLeaderState.Mutex.Unlock()

					if successfulReplications > nodeCount/2 {
						state.DefaultVolatileState.SetCommitIndex(CurrentLogIndex)
						lg.Log.Debugf("Replicated log on the majority of the nodes, new commit index: %d", CurrentLogIndex)
					}
				} else {
					if resp.Term > CurrentTerm {
						state.DefaultPersistentState.SetCurrentState(state.Follower)
					} else {
						state.DefaultLeaderState.Mutex.Lock()
						if state.DefaultLeaderState.NextIndex[clientIndex] > 1 {
							state.DefaultLeaderState.NextIndex[clientIndex]--
						}
						lg.Log.Debugf("Replicating log failed for node %d, retrying with NextIndex=%d", clientIndex, state.DefaultLeaderState.NextIndex[clientIndex])
						state.DefaultLeaderState.Mutex.Unlock()
					}
				}
			}(client, ctx, id)
		}

		time.Sleep(config.Default.HeartbeatInterval)
	}
}
