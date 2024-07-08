package tools

import (
	"errors"
	URL "net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

func NormalizeURL(url string) (string, error) {

	url = strings.ToLower(strings.TrimSpace(url))

	if len(url) == 0 {
		log.Error("Empty url")

		err := errors.New("empty url")

		return "", err
	}

	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	parsedURL, err := URL.Parse(url)
	if err != nil {

		log.Error("Could not parse url")
		return "", err

	}

	return parsedURL.Host + parsedURL.Path, nil
}
