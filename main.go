package main

import (
	"fmt"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	Example string
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "helloworld",
			Short:    "A Sensu Go \"hello world\" check plugin",
			Keyspace: "sensu.io/plugins/helloworld/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		&sensu.PluginConfigOption{
			Path:      "world",
			Env:       "SENSU_WORLD",
			Argument:  "world",
			Shorthand: "w",
			Default:   "",
			Usage:     "The world to send greetings from (optional)",
			Value:     &plugin.Example,
		},
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	if len(plugin.Example) > 0 {
		plugin.Example = fmt.Sprintf("%s ", plugin.Example) // add a trailing space
	}
	return sensu.CheckStateOK, nil
}

func executeCheck(event *types.Event) (int, error) {
	fmt.Printf("Hello, %sworld!\n", plugin.Example)
	return sensu.CheckStateOK, nil
}
