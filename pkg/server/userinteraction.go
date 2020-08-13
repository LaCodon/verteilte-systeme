package server

import (
	"context"
	"errors"
	"github.com/LaCodon/verteilte-systeme/internal/state"
	"github.com/LaCodon/verteilte-systeme/pkg/client"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
)

const (
	UserRequestSet    int32 = 1
	UserRequestDelete int32 = 2
	UserRequestGetAll int32 = 3
	UserRequestGetOne int32 = 4
)

func (s *Server) UserInteraction(c context.Context, ur *rpc.UserRequest) (*rpc.UserResponse, error) {
	// redirect to current leader
	if state.DefaultPersistentState.GetCurrentState() != state.Leader {
		return &rpc.UserResponse{
			ResponseCode: 301,
			RedirectTo:   state.DefaultVolatileState.GetCurrentLeader(),
		}, nil
	}

	//handle action
	if ur.RequestCode == UserRequestSet || ur.RequestCode == UserRequestDelete {
		//user wants to store a new value
		var action int32
		if ur.RequestCode == UserRequestSet {
			action = 1
		} else {
			action = 2
		}
		//send new input to channel
		client.NewUserInput <- &client.UserInput{
			Key:    ur.Key,
			Var:    ur.Value,
			Action: action,
		}

		return &rpc.UserResponse{
			ResponseCode: 202,
		}, nil
	} else if ur.RequestCode == UserRequestGetAll {
		// user wants all data
		return &rpc.UserResponse{
			ResponseCode: 200,
			Data:         state.DefaultPersistentState.GetCurrentStorage(),
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
			Data:         m,
		}, nil
	} else {
		// invalid request code
		return &rpc.UserResponse{
			ResponseCode: 400,
		}, errors.New("invalid RequestCode")
	}
}
