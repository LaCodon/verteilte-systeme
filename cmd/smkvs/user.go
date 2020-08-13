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
var actionString string

func init() {
	userInput = &client.UserInput{
		Action: 0,
		Key:    "",
		Var:    "",
	}
	nodeIp = "127.0.0.1:36000"
	actionString = ""
}

func setKVPair(c *cli.Context) error {
	userInput.Action = 1
	return sendRPC(server.UserRequestStore)
}

func deleteKVPair(c *cli.Context) error {
	userInput.Action = 2
	return sendRPC(server.UserRequestStore)
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
		Action:      userInput.Action,
		Key:         userInput.Key,
		Value:       userInput.Var,
	}

	for {
		// repeat request until responsecode is no redirect
		cs := client.ConnectToNodes([]string{nodeIp})
		ctx, _ := context.WithTimeout(context.Background(), config.Default.UserRequestTimeout)
		fmt.Printf("Sending request to %s\n", nodeIp)
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
