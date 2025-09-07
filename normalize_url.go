package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}

	host := parsedURL.Host
	path := parsedURL.EscapedPath()
	if host == "" {
		// No scheme, treat as already normalized
		host = parsedURL.Path
		path = ""
	}

	// Remove trailing slash unless path is just "/"
	if len(path) > 1 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	normalized := host + path

	if parsedURL.RawQuery != "" {
		normalized += "?" + parsedURL.RawQuery
	}
	if parsedURL.Fragment != "" {
		normalized += "#" + parsedURL.Fragment
	}

	return normalized, nil
}
