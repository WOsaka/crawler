package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s\n", args[0])

	html, err := getHTML(args[0])
	if err != nil {
		fmt.Printf("error fetching HTML: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(html)
}
