# Web Crawler

A concurrent web crawler written in Go that scrapes websites, extracts content, and generates a CSV report of the findings.

## Features

- **Concurrent Crawling**: Efficiently crawls websites using configurable concurrency limits
- **Domain Restriction**: Only crawls pages within the specified domain
- **Configurable Depth**: Limits crawling to a specified number of pages
- **Data Extraction**:
  - Page URL
  - H1 heading text
  - First paragraph text
  - Outgoing links
  - Image URLs
- **CSV Reporting**: Generates a comprehensive CSV report with all extracted data

## Installation

### Prerequisites

- Go 1.16 or higher

### Building from source

```bash
git clone https://github.com/WOsaka/crawler.git
cd crawler
go build
```

## Usage

```bash
./crawler <URL> <maxConcurrency> <maxPages>
```

### Parameters

- `URL`: The starting URL to begin crawling from
- `maxConcurrency`: The maximum number of concurrent requests allowed
- `maxPages`: The maximum number of pages to crawl

### Example

```bash
./crawler https://example.com 10 100
```

This command will:

1. Start crawling at `https://example.com`
2. Use a maximum of 10 concurrent goroutines for crawling
3. Crawl up to 100 pages within the example.com domain
4. Generate a report.csv file with the results

## Output

The crawler generates a CSV file named `report.csv` with the following columns:

- `page_url`: The URL of the crawled page
- `h1`: The H1 heading from the page
- `first_paragraph`: The first paragraph from the page
- `outgoing_link_urls`: Semicolon-separated list of outgoing links
- `image_urls`: Semicolon-separated list of image URLs

## Implementation Details

- Uses Go's concurrency primitives (goroutines, channels, and mutex) for efficient crawling
- Implements proper URL normalization to avoid duplicate crawling
- Respects robots.txt through a custom user agent
- Uses the goquery library for HTML parsing and content extraction
- Handles errors gracefully with appropriate logging

## Dependencies

- [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) - For HTML parsing
- Standard Go libraries

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
