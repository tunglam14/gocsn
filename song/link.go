package song

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetDownloadInfo(songURL string) map[string]string {
	pageURL := getDownloadPageURL(songURL)
	return getLinksFromDownloadPage(pageURL)
}

func getDownloadPageURL(songURL string) string {
	return strings.Replace(songURL, ".html", "_download.html", 1)
}

func getLinksFromDownloadPage(downloadPageURL string) map[string]string {
	doc, err := goquery.NewDocument(downloadPageURL)
	songDownloadLinks := make(map[string]string)

	if err != nil {
		fmt.Println("Parse download URL error")
		os.Exit(1)
	}

	doc.Find("#downloadlink a").Each(func(i int, s *goquery.Selection) {
		downloadLink, songBitrate, ok := parseDownloadLink(s)
		if ok {
			songDownloadLinks[songBitrate] = downloadLink
		}
	})

	return songDownloadLinks
}

func parseDownloadLink(s *goquery.Selection) (downLoadLink, songBitrate string, ok bool) {
	downloadLink, _ := s.Attr("href")
	if strings.Contains(downloadLink, "login") == true || downloadLink == "" {
		return "", "", false
	}

	re := regexp.MustCompile("\\[.*\\]")
	songBitrate = re.FindString(downloadLink)

	// remove [ and ] in bitrate
	songBitrate = strings.Replace(songBitrate, "[", "", -1)
	songBitrate = strings.Replace(songBitrate, "]", "", -1)

	// endcode URI
	downloadLink = urlEncode(downloadLink)

	return downloadLink, songBitrate, true
}

func urlEncode(unencoded string) string {
	var encoded *url.URL
	encoded, err := url.Parse(unencoded)
	if err != nil {
		return unencoded
	}
	return encoded.String()
}
