package main

import (
	"log"
	"testing"
)

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
				log.Println(actual)
				if actual != tc.expected {
					t.Errorf("expected %q, got %q", tc.expected, actual)
				}
			})
	}
}
