package main

import (
	"log"
	"net/url"
	"regexp"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	rawURL = strings.ToLower(rawURL)
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	host := parsedURL.Host
	host = strings.TrimPrefix(host, "www.")

	path := parsedURL.Path
	re := regexp.MustCompile("/+")
	path = re.ReplaceAllString(path, "/")

	normalizedURL := host + path
	log.Println(path)

	normalizedURL = strings.TrimSuffix(normalizedURL, "/")
	return normalizedURL, nil
}
