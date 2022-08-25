package main

import (
	core "github.com/gabbyprecious/demo-cln-plugin/pkg/plugin"

	"github.com/vincenzopalazzo/cln4go/plugin"
)

func main() {
	state := core.PluginState{}
	plugin := plugin.New(&state, true, plugin.DummyOnInit[core.PluginState])
	plugin.RegisterOption("foo", "string", "Hello Go", "An example of option", false)
	plugin.RegisterOption("demo-email", "string", "foo@email.com", "Email Plugin should send payment email to", false)

	plugin.RegisterRPCMethod("hello", "", "an example of rpc method", &core.Hello[core.PluginState]{})
	plugin.RegisterRPCMethod("networktester", "", "an example of rpc method", &core.NetworkChecker[core.PluginState]{})

	plugin.RegisterNotification("on-payment", &core.OnPayment[core.PluginState]{})
	plugin.RegisterNotification("shutdown", &core.OnShutdown[core.PluginState]{})
	plugin.Start()
}
