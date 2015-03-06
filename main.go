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
		///////////////////////////////////
		// i, err := strconv.ParseInt(args[1], 10, 64)
		// if err != nil {
		// 	fmt.Println("Num of route is invalid:", err)
		// 	os.Exit(1)
		// }
		// fmt.Println(i)

		// for x := 0; x < int(i); x++ {
		// 	// go cliConnection.CliCommand("files", "plugins")
		// 	go cliConnection.CliCommand("files", "plugins", "app/repo-index.yml")
		// }

		// go cliConnection.CliCommand("app", "cli")
		// c := make(chan int)
		// <-c
		///////////////////////////////////
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
			// cliConnection.CliCommand("map-route", args[1], args[3], "-n", args[4]+strconv.Itoa(x))
		}

	} else if args[0] == "del-routes" {
		if len(args) != 5 {
			fmt.Println("Invalid Usage: cf del-route DOMAIN HOST NUM_ROUTE OFFSET")
			os.Exit(1)
		}

		i, err := strconv.ParseInt(args[3], 10, 64)
		if err != nil {
			fmt.Println("Num of route is invalid:", err)
			os.Exit(1)
		}

		offset, err := strconv.ParseInt(args[4], 10, 64)
		if err != nil {
			fmt.Println("Num of offset is invalid:", err)
			os.Exit(1)
		}

		fmt.Printf("deleting routes ...\n\n")

		for x := int(offset); x < int(i+offset); x++ {
			cliConnection.CliCommand("delete-route", args[1], "-n", args[2]+strconv.Itoa(x), "-f")
		}

	} else if args[0] == "get-apps" {
		if len(args) != 2 {
			fmt.Println("Invalid Usage: cf get-apps NUM_OFF_RUN")
			os.Exit(1)
		}

		i, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Println("Num of run is invalid:", err)
			os.Exit(1)
		}

		for x := 0; x < int(i); x++ {
			fmt.Println("RUN #", x)
			cliConnection.CliCommand("apps")
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
					Usage: "cf gen-route APP_NAME SPACE DOMAIN HOST NUM_ROUTE OFFSET",
				},
			},
			{
				Name:     "del-routes",
				Alias:    "dr",
				HelpText: "Delete generated routes for an app",
				UsageDetails: plugin.Usage{
					Usage: "cf del-route DOMAIN HOST NUM_ROUTE OFFSET",
				},
			},
			{
				Name:     "get-apps",
				Alias:    "ga",
				HelpText: "repeatedly run `apps` command",
				UsageDetails: plugin.Usage{
					Usage: "cf get-apps NUM_OFF_RUN",
				},
			},
		},
	}
}
