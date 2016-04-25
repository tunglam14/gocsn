package song

import "fmt"
import "github.com/codegangsta/cli"

const (
	csnDomain    = "http://chiasenhac.com"
	searchURL    = "http://search.chiasenhac.com/search.php?s="
	searchResult = 2
)

type Song struct {
	Name          string
	Quanlity      string
	DownloadTotal string
	URL           string
	DownloadLink  map[string]string
}

func (song Song) Show() {
	fmt.Printf("%s | %s | Download: %s \nSong play link: %s \nDownload link:\n", song.Name, song.Quanlity, song.DownloadTotal, song.URL)
	for bitrate, link := range song.DownloadLink {
		fmt.Printf("%s: %s\n", bitrate, link)
	}
	fmt.Printf("\n\n")
}

func GetNameFromArgs(args cli.Args) string {
	argsCapacity := cap(args)
	var songName string
	for i := 0; i < argsCapacity; i++ {
		songName += args[i]
		if i < (argsCapacity - 1) {
			songName += " "
		}
	}
	return songName
}

func ShowSearchResult(songName string) {
	songs := SearchByName(songName)

	for _, song := range songs {
		song.Show()
	}

	return
}
