package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/tunglam14/gocsn/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "search",
		Usage:  "Search a song. Usage: csn search <song_name>",
		Action: command.CmdSearch,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "get",
		Usage:  "Download a song. Usage: csn get <song_id>",
		Action: command.CmdGet,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.\n", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
