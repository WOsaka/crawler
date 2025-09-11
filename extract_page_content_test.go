package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name      string
		inputHtml string
		pageURL   string
		expected  PageData
	}{{
		name:    "basic test",
		pageURL: "https://blog.boot.dev",
		inputHtml: "<html><body>" +
			"<h1>Test Title</h1>" +
			"<p>This is the first paragraph.</p>" +
			"<a href=\"/link1\">Link 1</a>" +
			"<img src=\"/image1.jpg\" alt=\"Image 1\">" +
			"</body></html>",
		expected: PageData{
			URL:            "https://blog.boot.dev",
			H1:             "Test Title",
			FirstParagraph: "This is the first paragraph.",
			OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
			ImageURLs:      []string{"https://blog.boot.dev/image1.jpg"},
		},
	},
		{

			name:    "multiple images and links",
			pageURL: "https://blog.boot.dev",
			inputHtml: "<html><body>" +
				"<h1>Test Title</h1>" +
				"<p>This is the first paragraph.</p>" +
				"<img src=\"/image1.jpg\" alt=\"Image 1\">" +
				"<a href=\"/link1\">Link 1</a>" +
				"<img src=\"/image2.jpg\" alt=\"Image 2\">" +
				"<a href=\"/link2\">Link 2</a>" +
				"</body></html>",
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  []string{"https://blog.boot.dev/link1", "https://blog.boot.dev/link2"},
				ImageURLs:      []string{"https://blog.boot.dev/image1.jpg", "https://blog.boot.dev/image2.jpg"},
			},
		},
		{
			name:    "no images or links",
			pageURL: "https://blog.boot.dev",
			inputHtml: "<html><body>" +
				"<h1>Test Title</h1>" +
				"<p>This is the first paragraph.</p>" +
				"</body></html>",
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  []string{},
				ImageURLs:      []string{},
			},
		},
		{
			name:    "no h1 or paragraph",
			pageURL: "https://blog.boot.dev",
			inputHtml: "<html><body>" +
				"<a href=\"/link1\">Link 1</a>" +
				"<img src=\"/image1.jpg\" alt=\"Image 1\">" +
				"</body></html>",
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "",
				FirstParagraph: "",
				OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
				ImageURLs:      []string{"https://blog.boot.dev/image1.jpg"},
			},
		},
		{
			name:      "empty HTML",
			pageURL:   "https://blog.boot.dev",
			inputHtml: "",
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "",
				FirstParagraph: "",
				OutgoingLinks:  []string{},
				ImageURLs:      []string{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractPageData(tc.inputHtml, tc.pageURL)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{{
		name:     "no images",
		inputURL: "https://blog.boot.dev",
		inputBody: "<html><body>" +
			"<a href=\"https://other.com/path/one\">" +
			"<span>Boot.dev</span>" +
			"</a>" +
			"</body></html>",
		expected: []string{},
	},
		{
			name:     "relative images",
			inputURL: "https://blog.boot.dev",
			inputBody: "<html><body>" +
				"<img src=\"/path/to/image.png\"/>" +
				"</body></html>",
			expected: []string{"https://blog.boot.dev/path/to/image.png"},
		},
		{
			name:     "absolute images",
			inputURL: "https://blog.boot.dev",
			inputBody: "<html><body>" +
				"<img src=\"https://other.com/path/to/image.png\"/>" +
				"</body></html>",
			expected: []string{"https://other.com/path/to/image.png"},
		},
		{
			name:     "find all <img> URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: "<html><body>" +
				"<img src=\"/path/to/image1.png\"/>" +
				"<img src=\"https://other.com/path/to/image2.png\"/>" +
				"</body></html>",
			expected: []string{
				"https://blog.boot.dev/path/to/image1.png",
				"https://other.com/path/to/image2.png",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("couldn't parse input URL: %v", err)
				return
			}

			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			expected := tc.expected
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("expected %v, got %v", expected, actual)
			}
		})
	}
}

func TestGetURLFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{{
		name:     "absolute URLs",
		inputURL: "https://blog.boot.dev",
		inputBody: "<html><body>" +
			"<a href=\"https://other.com/path/one\">" +
			"<span>Boot.dev</span>" +
			"</a>" +
			"</body></html>",
		expected: []string{"https://other.com/path/one"},
	},
		{
			name:     "relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: "<html><body>" +
				"<a href=\"/path/one\">" +
				"<span>Boot.dev</span>" +
				"</a>" +
				"</body></html>",
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "no URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: "<html><body>" +
				"<span>Boot.dev</span>" +
				"</body></html>",
			expected: []string{},
		},
		{
			name:     "find all <a> URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: "<html><body>" +
				"<a href=\"/path/one\">" +
				"<span>Boot.dev</span>" +
				"</a>" +
				"<a href=\"https://other.com/path/two\">" +
				"<span>Boot.dev</span>" +
				"</a>" +
				"</body></html>",
			expected: []string{
				"https://blog.boot.dev/path/one",
				"https://other.com/path/two",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("couldn't parse input URL: %v", err)
				return
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			expected := tc.expected
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("expected %v, got %v", expected, actual)
			}
		})
	}
}

func TestGetH1FromHTML(t *testing.T) {
	tests := []struct {
		inputBody string
		expected  string
	}{{
		inputBody: "<html><body><h1>Test Title</h1></body></html>",
		expected:  "Test Title",
	},
		{
			inputBody: "<html><body></body></html>",
			expected:  "",
		},
		{inputBody: "<html><body><h1>First Title</h1><h1>Second Title</h1></body></html>",
			expected: "First Title",
		},
	}

	for i, tc := range tests {
		t.Run(
			string(rune(i+'0')),
			func(t *testing.T) {
				actual := getH1FromHTML(tc.inputBody)

				if actual != tc.expected {
					t.Errorf("expected %q, got %q", tc.expected, actual)
				}
			})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	tests := []struct {
		inputBody string
		expected  string
	}{
		{
			inputBody: "<html><body>" +
				"<p>Outside paragraph.</p>" +
				"<main>" +
				"<p>Main paragraph.</p>" +
				"<p>Second paragraph.</p>" +
				"<main>" +
				"</body></html>",
			expected: "Main paragraph."},
		{
			inputBody: "<html><body>" +
				"<p>Main paragraph.</p>" +
				"</body></html>",
			expected: "Main paragraph.",
		},
		{
			inputBody: "<html><body>" +
				"</body></html>",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(
			string(rune(i+'0')),
			func(t *testing.T) {
				actual := getFirstParagraphFromHTML(tc.inputBody)
				if actual != tc.expected {
					t.Errorf("expected %q, got %q", tc.expected, actual)
				}
			})
	}
}
