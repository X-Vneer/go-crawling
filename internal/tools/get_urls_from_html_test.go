package tools

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{"<a href=\"https://example1.com/\">Example</a>", []string{}},
		{"<a href=\"#something\">Example</a>", []string{}},
		{"<a href=\"/path/path2/\">Example</a>", []string{"https://example.com/path/path2"}},
		{"<a href=\"https://example.com/path/path2/\">Example</a>", []string{"https://example.com/path/path2"}},
		{"<a href=\"https://example1.com/\">Example</a><a href=\"/path/path2/\">Example</a>", []string{"https://example.com/path/path2"}},
		{"<a href=\"/path/path2/\">Example</a><a href=\"https://example.com/\">Example</a>", []string{"https://example.com/path/path2", "https://example.com"}},
		{"<a href=\"/path/path2/\">Example</a><a href=\"/path/path3/\">Example</a>", []string{"https://example.com/path/path2", "https://example.com/path/path3"}},
	}

	for index, c := range cases {
		got := GetURLsFromHTML(c.input, "https://example.com/")
		fmt.Printf("testing GetURLsFromHTML case #%d: %v  \n\n", index, reflect.DeepEqual(got, c.expected) || len(got) == len(c.expected) && len(got) == 0)
		if len(got) == len(c.expected) && len(got) == 0 {
			continue
		}
		if len(got) != len(c.expected) {
			t.Errorf("GetURLsFromHTML(%s) = %v, expected %v", c.input, got, c.expected)
		}
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("GetURLsFromHTML(%s) = %v, expected %v", c.input, got, c.expected)
		}
	}
}
