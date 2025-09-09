package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return ""
	}
	h1Tag := doc.Find("h1").First().Text()
	return h1Tag
}

func getFirstParagraphFromHTML(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return ""
	}
	pTag := doc.Find("main").Find("p").First().Text()
	if pTag == "" {
		pTag = doc.Find("p").First().Text()
	}
	return pTag 
}
