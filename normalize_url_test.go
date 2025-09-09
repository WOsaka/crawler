package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing /",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove www",
			inputURL: "https://www.blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "no changes needed",
			inputURL: "blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "lowercase hostname",
			inputURL: "http://EXAMPLE.COM/path",
			expected: "example.com/path",
		},
		{
			name:     "URL with port",
			inputURL: "https://example.com:8080/path",
			expected: "example.com:8080/path",
		},
		{
			name:     "URL with fragment",
			inputURL: "https://example.com/path#section",
			expected: "example.com/path",
		},
		{
			name:     "multiple slashes",
			inputURL: "https://example.com//foo///bar/",
			expected: "example.com/foo/bar",
		},
		{
			name:     "uppercase host and path",
			inputURL: "https://EXAMPLE.COM/FOO/BAR",
			expected: "example.com/foo/bar",
		},
		{
			name:     "user info in URL",
			inputURL: "https://user:pass@example.com/path",
			expected: "example.com/path",
		},
		{
			name:     "empty path",
			inputURL: "https://example.com",
			expected: "example.com",
		},
		{
			name:     "only slash path",
			inputURL: "https://example.com/",
			expected: "example.com",
		},
		{
			name:     "subdomain",
			inputURL: "https://sub.example.com/path",
			expected: "sub.example.com/path",
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
