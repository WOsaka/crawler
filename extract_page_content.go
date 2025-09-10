package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	imgs := []string{}
	doc.Find("img[src]").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		if len(src) > 0 && src[0] == '/' {
			baseURL.Path = src
			imgs = append(imgs, baseURL.String())
		} else {
			imgs = append(imgs, src)
		}
	})

	return imgs, nil
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	urls := []string{}
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if len(href) > 0 && href[0] == '/' {
			baseURL.Path = href
			urls = append(urls, baseURL.String())
		} else {
			urls = append(urls, href)
		}
	})

	return urls, nil
}

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
	// pTag := doc.Find("main").Find("p").First().Text()
	pTag := doc.Find("main p").First().Text()
	if pTag == "" {
		pTag = doc.Find("p").First().Text()
	}
	return pTag
}
