package plugin

import (
	"os"
	"testing"

	"github.com/vincenzopalazzo/cln4go/client"
)

func TestCallFistMethod(t *testing.T) {
	path := os.Getenv("CLN_UNIX_SOCKET")
	client, err := client.NewUnix(path)
	if err != nil {
		panic(err)
	}
	response, err := client.Call("hello", make(map[string]interface{}))
	if err != nil {
		panic(err)
	}

	message, found := response["message"]
	if !found {
		t.Error("The message is not found")
	}

	if message != "hello from cln4go.template" {
		t.Errorf("message received %s different from expected %s", message, "hello from cln4go.template")
	}
}

func TestCallNetworkMethod(t *testing.T) {
	path := os.Getenv("CLN_UNIX_SOCKET")
	client, err := client.NewUnix(path)
	if err != nil {
		panic(err)
	}
	response, err := client.Call("networktester", make(map[string]interface{}))
	if err != nil {
		panic(err)
	}

	message, found := response["message"]
	if !found {
		t.Error("The message is not found")
	}
	network := os.Getenv("NETWORK")
	if message != "running on "+network {
		t.Errorf("message received %s different from expected %s", message, "running on "+network)
	}
}
