package main

import (
	"bufio"
	"fmt"
	URL "net/url"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/x-vneer/crawling/internal/crawling"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("URL to be crawled: ")
	url, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Could not read a valid url")
	}

	url = strings.TrimSpace(url)

	parsedURL, err := URL.Parse(url)
	if err != nil {
		log.Fatal("Could not parse url")
	}

	var crawledPages map[string]uint8 = make(map[string]uint8)
	crawling.Crawling(parsedURL.Scheme+"://"+parsedURL.Host, url, &crawledPages)

	if err := os.MkdirAll("results", 0755); err != nil {
		log.Fatal("failed to create directory")
	}

	result := ""
	for k, v := range crawledPages {
		result += fmt.Sprintf("%-10v %s\n", fmt.Sprintf("%d:", v), k)
	}

	os.WriteFile("results/"+strings.Replace(parsedURL.Host, ".", "_", -1)+".txt", []byte(result), 0644)

}
