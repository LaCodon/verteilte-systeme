package main

import (
	"context"
	"fmt"
	"github.com/LaCodon/verteilte-systeme/pkg/client"
	"github.com/LaCodon/verteilte-systeme/pkg/config"
	"github.com/LaCodon/verteilte-systeme/pkg/rpc"
	"github.com/LaCodon/verteilte-systeme/pkg/server"
	"github.com/urfave/cli/v2"
)

var userInput *client.UserInput
var nodeIp string

func init() {
	userInput = &client.UserInput{
		Action: 0,
		Key:    "",
		Var:    "",
	}
	nodeIp = "127.0.0.1:36000"
}

func setKVPair(c *cli.Context) error {
	return sendRPC(server.UserRequestSet)
}

func deleteKVPair(c *cli.Context) error {
	return sendRPC(server.UserRequestDelete)
}

func getStorage(c *cli.Context) error {
	if userInput.Key != "" {
		return sendRPC(server.UserRequestGetOne)
	} else {
		return sendRPC(server.UserRequestGetAll)
	}
}

func sendRPC(requestCode int32) error {
	req := &rpc.UserRequest{
		RequestCode: requestCode,
		Key:         userInput.Key,
		Value:       userInput.Var,
	}

	for {
		// repeat request until responsecode is no redirect
		cs := client.ConnectToNodes([]string{nodeIp})
		ctx, _ := context.WithTimeout(context.Background(), config.Default.UserRequestTimeout)
		resp, err := cs[0].NodeClient.UserInteraction(ctx, req)
		if err != nil {
			return err
		}
		if resp.ResponseCode == 301 {
			nodeIp = resp.RedirectTo
			continue
		}
		// request sent to leader -> end of loop
		if resp.ResponseCode == 202 {
			fmt.Print("Input accepted\n")
		} else if resp.Data != nil {
			fmt.Printf("Data: %+v\n", resp.Data)
		}
		return nil
	}

}
