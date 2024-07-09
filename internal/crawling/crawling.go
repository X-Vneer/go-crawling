package crawling

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/x-vneer/crawling/internal/tools"
)

func Crawling(baseURL string, currentURL string, crawledPages *map[string]uint8) {

	fmt.Printf("\nCrawling: %s\n", currentURL)

	url := tools.NormalizeURL(currentURL)

	if len(url) == 0 {
		log.Warn("Empty url")
		return
	}

	_, ok := (*crawledPages)[url]

	(*crawledPages)[url] += 1

	if ok {
		return
	}

	page, err := tools.GetHTMLPage(currentURL)
	if err != nil {
		log.Error(err)
		return
	}

	// Get all URLs from currentURL
	urls := tools.GetURLsFromHTML(page, baseURL)

	// Crawl each URL
	for _, url := range urls {
		Crawling(baseURL, url, crawledPages)
	}

}
