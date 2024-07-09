package main

import (
	"fmt"

	"github.com/x-vneer/crawling/internal/crawling"
)

func main() {

	// TODO: Implement crawling

	var crawledPages map[string]uint8 = make(map[string]uint8)
	crawling.Crawling("https://mabet.com.sa", "https://mabet.com.sa/ar", &crawledPages)

	fmt.Println(crawledPages)
}
