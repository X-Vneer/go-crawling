package tools

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

func GetURLsFromHTML(html string, baseURL string) []string {
	// Parse the HTML string
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	// Find the URL
	var urls []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		if !exists {
			return
		}

		if href[0] == '/' || href[0] == '.' {
			result, _ := url.JoinPath(baseURL, href)
			result = removeTrailingSlash(result)
			urls = append(urls, result)
		}

		if strings.HasPrefix(href, baseURL) {
			result, _ := url.Parse(href)
			urlString := removeTrailingSlash(result.String())
			urls = append(urls, urlString)

		}

	})

	return urls
}

// RemoveTrailingSlash remove trailing slash from url if exists
func removeTrailingSlash(url string) string {
	if url[len(url)-1] == '/' {
		return url[:len(url)-1]
	}
	return url
}
