package song

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func SearchByName(songName string) []Song {
	song := findSongInfo(songName)
	return song
}

func findSongInfo(songName string) []Song {
	songNameQueryString := strings.Replace(songName, " ", "+", -1)
	songSearchURL := fmt.Sprintf("%s%s", searchURL, songNameQueryString)

	doc, err := goquery.NewDocument(songSearchURL)
	if err != nil {
		fmt.Println("Parse song URL error")
		os.Exit(1)
	}

	// var songs []Song
	songs := make([]Song, searchResult)

	doc.Find(".bod .tbtable tr").Each(func(i int, s *goquery.Selection) {
		// i=0 means head of table (title), loop should start with i = 1
		if i > 0 && i <= searchResult {
			c := make(chan Song)
			go getSongData(s, c)

			// loop start from 1
			// slide start from 0
			songs[i-1] = <-c
		}
	})

	return songs
}

func getSongData(s *goquery.Selection, c chan Song) {
	var song Song
	s.Find("td").Each(func(i int, songLine *goquery.Selection) {
		lineValue := formatSongData(songLine.Text())
		switch i {
		case 1:
			song.Name = lineValue
			song.URL = getSongURL(songLine)
		case 2:
			song.Quanlity = lineValue
		case 5:
			song.DownloadTotal = lineValue
		}
	})
	song.DownloadLink = GetDownloadInfo(song.URL)

	c <- song
}

func getSongURL(s *goquery.Selection) string {
	songURL, _check := s.Find(".musictitle").Attr("href")

	if _check {
		return fmt.Sprintf("%s/%s", csnDomain, songURL)
	}

	return "URL not found"
}

// Remove whitespace, tab, newline chars
func formatSongData(in string) (out string) {
	out = strings.TrimSpace(in)
	out = strings.Replace(out, "\t", "", -1)
	out = strings.Replace(out, "\n", " - ", -1)
	return out
}
