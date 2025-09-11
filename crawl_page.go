package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Printf("error parsing base URL %s: %v", rawBaseURL, err)
		return
	}

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("error parsing current URL %s: %v", rawBaseURL, err)
		return
	}

	if parsedBaseURL.Host != parsedCurrentURL.Host {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("error normalizing URL %s: %v", rawCurrentURL, err)
		return
	}

	if _, found := pages[normalizedCurrentURL]; found {
		pages[normalizedCurrentURL]++
		return
	} else {
		pages[normalizedCurrentURL] = 1
	}

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("error fetching HTML for %s: %v", rawCurrentURL, err)
		return
	}
	// log.Println(html)

	links, err := getURLsFromHTML(html, parsedBaseURL)
	if err != nil {
		log.Printf("error extracting links from %s: %v", rawCurrentURL, err)
		return
	}
	log.Printf("found %d links on %s\n", len(links), rawCurrentURL)

	for _, link := range links {
		crawlPage(rawBaseURL, link, pages)
	}

	if rawBaseURL == rawCurrentURL {
		fmt.Println("Crawl complete, visited pages:")
		fmt.Println(pages) 
	}
}
