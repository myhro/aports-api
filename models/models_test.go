package models

import (
	"net/url"
	"testing"
)

func TestGetOffset(t *testing.T) {
	table := []struct {
		in   string
		want int
	}{
		{"?nopage=", 0},
		{"?page=invalid_number", 0},
		{"?page=", 0},
		{"?page=0", 0},
		{"?page=1", 0},
		{"?page=2", ResultsPerPage},
		{"?page=3", ResultsPerPage * 2},
	}

	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			u, err := url.Parse(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			got := getOffset(u.Query())
			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceWildcards(t *testing.T) {
	table := []struct {
		in   string
		want string
	}{
		{"bash*", "bash%"},
		{"*bash*", "%bash%"},
		{"bas?", "bas_"},
		{"?as?", "_as_"},
	}

	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			got := replaceWildcards(tt.in)
			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
