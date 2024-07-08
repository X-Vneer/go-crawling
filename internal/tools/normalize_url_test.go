package tools

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"http://example.com/", "example.com"},
		{"http://exaMpLe.cOm/", "example.com"},
		{"http://example.com/path", "example.com/path"},
		{"http://example.com/path/", "example.com/path"},
		{"http://example.com/path/path2", "example.com/path/path2"},
		{"http://example.com/path/path2/", "example.com/path/path2"},
		{"http://example.com/path/path2/path3", "example.com/path/path2/path3"},
		{"https://example.com/", "example.com"},
		{"https://example.com/path", "example.com/path"},
		{"https://example.com/path/", "example.com/path"},
		{"https://example.com/path/path2", "example.com/path/path2"},
		{"https://example.com/path/path2/", "example.com/path/path2"},
		{"https://example.com/path/path2/path3", "example.com/path/path2/path3"},
	}

	for _, c := range cases {
		normalizedURL, err := NormalizeURL(c.input)
		if err != nil {
			t.Errorf("NormalizeURL(%s) returned error: %v", c.input, err)
		}
		if normalizedURL != c.expected {
			t.Errorf("NormalizeURL(%s) = %s, expected %s", c.input, normalizedURL, c.expected)
		}
	}

}
