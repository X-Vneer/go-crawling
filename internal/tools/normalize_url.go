package tools

import (
	URL "net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

func NormalizeURL(url string) string {

	url = strings.ToLower(strings.TrimSpace(url))

	if len(url) == 0 {
		log.Error("Empty url")
		return ""
	}

	parsedURL, err := URL.Parse(url)

	if err != nil {
		log.Error("Could not parse url")
		return ""
	}

	parsedURL.Path = strings.Split(parsedURL.Path, "?")[0]

	if len(parsedURL.Path) > 0 && parsedURL.Path[len(parsedURL.Path)-1] == '/' {
		parsedURL.Path = parsedURL.Path[:len(parsedURL.Path)-1]
	}

	return parsedURL.Host + parsedURL.Path
}
