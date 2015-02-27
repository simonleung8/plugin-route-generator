/**
* This is an example plugin where we use both arguments and flags. The plugin
* will echo all arguments passed to it. The flag -uppercase will upcase the
* arguments passed to the command.
**/
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cloudfoundry/cli/plugin"
)

type GenRoutes struct {
}

func main() {
	plugin.Start(new(GenRoutes))
}

func (g *GenRoutes) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "gen-routes" {
		if len(args) != 7 {
			fmt.Println("Invalid Usage: cf gen-route APP_NAME SPACE DOMAIN HOST NUM_ROUTE OFFSET")
			os.Exit(1)
		}

		i, err := strconv.ParseInt(args[5], 10, 64)
		if err != nil {
			fmt.Println("Num of route is invalid:", err)
			os.Exit(1)
		}
		offset, err := strconv.ParseInt(args[6], 10, 64)
		if err != nil {
			fmt.Println("Num of offset is invalid:", err)
			os.Exit(1)
		}

		fmt.Printf("Generating %d routes for '%s'...\n\n", i, args[2])

		for x := int(offset); x < int(i+offset); x++ {
			cliConnection.CliCommand("create-route", args[2], args[3], "-n", args[4]+strconv.Itoa(x))
			cliConnection.CliCommand("map-route", args[1], args[3], "-n", args[4]+strconv.Itoa(x))
		}

	} else if args[0] == "del-routes" {
		if len(args) != 4 {
			fmt.Println("Invalid Usage: cf del-route APP_NAME SPACE DOMAIN HOST NUM_ROUTE")
			os.Exit(1)
		}

		i, err := strconv.ParseInt(args[3], 10, 64)
		if err != nil {
			fmt.Println("Num of route is invalid:", err)
			os.Exit(1)
		}

		fmt.Printf("deleting routes ...\n\n")

		for x := 0; x < int(i); x++ {
			cliConnection.CliCommand("delete-route", args[1], "-n", args[2]+strconv.Itoa(x), "-f")
		}

	}
}

func (g *GenRoutes) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Route-Generator",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "gen-routes",
				Alias:    "gr",
				HelpText: "Generate routes for an app",
				UsageDetails: plugin.Usage{
					Usage: "cf gen-route APP_NAME SPACE DOMAIN HOST NUM_ROUTE",
				},
			},
			{
				Name:     "del-routes",
				Alias:    "dr",
				HelpText: "Delete generated routes for an app",
				UsageDetails: plugin.Usage{
					Usage: "cf del-route DOMAIN HOST NUM_ROUTE",
				},
			},
		},
	}
}
