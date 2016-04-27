package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/tunglam14/gocsn/song"
)

func CmdSearch(c *cli.Context) {
	songName := c.Args()[0]
	fmt.Printf("Search: %s \n=======", songName)
	song.ListByName(songName)
}
