package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	headers := []string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"}
	if err := w.Write(headers); err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	for _, page := range pages {
		outgoingLinks := strings.Join(page.OutgoingLinks, ";")
		imageUrls := strings.Join(page.ImageURLs, ";")
		record := []string{page.URL, page.H1, page.FirstParagraph, outgoingLinks, imageUrls}
		w.Write(record)
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return fmt.Errorf("error writing CSV: %w", err)
	}
	return nil
}
