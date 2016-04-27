package song

import "fmt"

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

func (song Song) show() {
	fmt.Printf("%s | %s | Download: %s \nSong play link: %s \nDownload link:\n", song.Name, song.Quanlity, song.DownloadTotal, song.URL)
	for bitrate, link := range song.DownloadLink {
		fmt.Printf("%s: %s\n", bitrate, link)
	}
	fmt.Printf("\n\n")
}
