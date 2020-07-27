package server

import (
	"context"
	"errors"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/client"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
)

const (
	UserRequestStore int32 = 1
	UserRequestGetAll int32 = 2
	UserRequestGetOne int32 = 3
)

func (s *Server) UserInteraction(c context.Context, ur *rpc.UserRequest) (*rpc.UserResponse, error) {
	if ur.RequestCode == UserRequestStore {
		//user wants to store a new value

		// check for valid action
		if ur.Action < 1 || ur.Action > 2{
			return &rpc.UserResponse{
				ResponseCode: 400,
			}, errors.New("invalid action")
		}

		// redirect to current leader
		if state.DefaultPersistentState.GetCurrentState() != state.Leader {
			return &rpc.UserResponse{
				ResponseCode: 301,
				RedirectTo: state.DefaultVolatileState.GetCurrentLeader(),
			}, nil
		}

		//send new input to channel
		client.NewUserInput <- &client.UserInput{
			Key:    ur.Key,
			Var:    ur.Value,
			Action: ur.Action,
		}

		return &rpc.UserResponse{
			ResponseCode: 202,
		}, nil
	} else if ur.RequestCode == UserRequestGetAll {
		// user wants all data
		return &rpc.UserResponse{
			ResponseCode: 200,
			Data: state.DefaultPersistentState.GetCurrentStorage(),
		}, nil
	} else if ur.RequestCode == UserRequestGetOne {
		// user wants only current value of key
		var m map[string]string
		value, ok := state.DefaultPersistentState.GetCurrentValue(ur.Key)
		if ok {
			m := make(map[string]string)
			m[ur.Key] = value
		}
		return &rpc.UserResponse{
			ResponseCode: 200,
			Data: m,
		}, nil
	} else {
		// invalid request code
		return &rpc.UserResponse{
			ResponseCode: 400,
		}, errors.New("invalid RequestCode")
	}
}

