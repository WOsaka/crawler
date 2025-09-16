package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func handleArgs() (*url.URL, int, int) {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	} else if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	parsedURL, err := url.Parse(args[0])
	if err != nil {
		fmt.Printf("invalid URL: %v\n", err)
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil || maxConcurrency <= 0 {
		fmt.Println("max_concurrency must be a positive integer")
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil || maxPages <= 0 {
		fmt.Println("max_pages must be a positive integer")
		os.Exit(1)
	}

	return parsedURL, maxConcurrency, maxPages
}
