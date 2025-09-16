package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	baseURL, maxConcurrency, maxPages := handleArgs()

	cfg := &config{
		pages:              make(map[string]PageData),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	fmt.Printf("starting crawl of: %s\n", baseURL.String())
	cfg.crawlPage(os.Args[1])

	cfg.wg.Wait()
}
