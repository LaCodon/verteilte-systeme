package server

import (
	"context"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
)

func (s *Server) RequestVote(context.Context, *rpc.VoteRequest) (*rpc.VoteResponse, error) {
	return &rpc.VoteResponse{}, nil
}
