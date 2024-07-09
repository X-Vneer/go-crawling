package tests

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/x-vneer/crawling/internal/tools"
)

func TestNormalizeURL(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"http://example.com/", "example.com"},
		{"https://example.com/", "example.com"},
		{"http://exaMpLe.cOm/", "example.com"},
		{"http://example.com/path", "example.com/path"},
		{"http://example.com/path/", "example.com/path"},
		{"https://example.com/path?search=query", "example.com/path"},
		{"https://mabet.com.sa/ar/units/%3Fq=camps&type=4", "mabet.com.sa/ar/units"},
		{"", ""},
	}

	for index, c := range cases {
		normalizedURL := tools.NormalizeURL(c.input)
		fmt.Printf("testing NormalizeURL case #%-10v %v  \n\n", strconv.Itoa(index)+":", (normalizedURL == c.expected))

		if normalizedURL != c.expected {
			t.Errorf("NormalizeURL(%s) = %s, expected %s", c.input, normalizedURL, c.expected)
		}
	}

}
