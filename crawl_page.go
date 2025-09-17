package main

import (
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.mu.Lock()
	pagesLength := len(cfg.pages)
	if pagesLength >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	// linkStatus := "not followed"
	// defer func() {
	// 	log.Printf("%s - %s", rawCurrentURL, linkStatus)
	// }()

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("error parsing current URL %s: %v", rawCurrentURL, err)
		return
	}

	if cfg.baseURL.Host != parsedCurrentURL.Host {
		// linkStatus = "not followed - external domain"
		return
	}

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("error fetching HTML for %s: %v", rawCurrentURL, err)
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("error normalizing URL %s: %v", rawCurrentURL, err)
		return
	}

	if !cfg.firstVisit(normalizedCurrentURL) {
		// linkStatus = "not followed - already visited"
		return
	}

	// linkStatus = "followed"

	pageData := extractPageData(html, rawCurrentURL)

	cfg.mu.Lock()
	cfg.pages[normalizedCurrentURL] = pageData
	cfg.mu.Unlock()
	// log.Println("added to pages:", normalizedCurrentURL)

	for _, link := range pageData.OutgoingLinks {
		cfg.wg.Add(1)
		go func(url string) {
			defer cfg.wg.Done()
			cfg.concurrencyControl <- struct{}{}
			cfg.crawlPage(url)
			<-cfg.concurrencyControl
		}(link)
	}

}

func (cfg *config) firstVisit(normalizedURL string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, found := cfg.pages[normalizedURL]
	if found {
		return false
	}

	cfg.pages[normalizedURL] = PageData{}
	return true
}
