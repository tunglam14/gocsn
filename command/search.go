package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/tunglam14/csn/song"
)

func CmdSearch(c *cli.Context) {
	songName := song.GetNameFromArgs(c.Args())
	fmt.Println("Search:", songName)
	fmt.Println("=======")
	song.ShowSearchResult(songName)
}
