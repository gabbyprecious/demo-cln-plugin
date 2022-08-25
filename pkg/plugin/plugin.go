package plugin

import (
	"fmt"
	"os"
	"strings"

	"github.com/vincenzopalazzo/cln4go/client"

	"github.com/vincenzopalazzo/cln4go/plugin"
)

func sendMail(email string) {
	// send mail or do something else
}

type PluginState struct{}

type Hello[T PluginState] struct{}

func (instance *Hello[T]) Call(plugin *plugin.Plugin[T], request map[string]any) (map[string]any, error) {
	return map[string]any{"message": "hello from cln4go.template"}, nil
}

type NetworkChecker[T PluginState] struct{}

func (instance *NetworkChecker[PluginState]) Call(plugin *plugin.Plugin[PluginState], request map[string]any) (map[string]any, error) {
	clnPath, found := plugin.GetConf("lightning-dir")
	if !found {
		panic(found)
	}

	rpcFileName, found := plugin.GetConf("rpc-file")
	if !found {
		panic(found)
	}
	unixPath := strings.Join([]string{clnPath.(string), rpcFileName.(string)}, "/")
	client, err := client.NewUnix(unixPath)
	if err != nil {
		return map[string]any{"message": "error setting up client"}, err
	}
	response, err := client.Call("getinfo", make(map[string]interface{}))

	if err != nil {
		return map[string]any{"message": "error calling getinfo"}, err
	}
	network := response["network"]
	if !found {
		return map[string]any{"message": "unknown"}, nil
	}
	return map[string]any{"message": fmt.Sprintf("running on %s", network.(string))}, nil
}

type OnPayment[T PluginState] struct{}

func (instance *OnPayment[PluginState]) Call(plugin *plugin.Plugin[PluginState], request map[string]any) {
	receivingEmail, found := plugin.GetOpt("demo-email")

	if !found {
		receivingEmail = "bar@email.com"
	}
	sendMail(receivingEmail.(string))
}

type OnShutdown[T PluginState] struct{}

func (instance *OnShutdown[T]) Call(plugin *plugin.Plugin[T], request map[string]any) {
	os.Exit(0)
}
