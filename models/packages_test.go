package models

import (
	"net/url"
	"testing"
)

func TestPackagesWhereClause(t *testing.T) {
	table := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"?invalid_field=1", ""},
		{"?name=bash", "WHERE packages.name = :packages.name"},
		{"?name=bash&repo=", "WHERE packages.name = :packages.name"},
		{"?name=bash&repo=&arch=", "WHERE packages.name = :packages.name"},
		{"?name=bash&repo=main", "WHERE packages.name = :packages.name AND packages.repo = :packages.repo"},
		{"?name=bash&repo=main&arch=x86", "WHERE packages.arch = :packages.arch AND packages.name = :packages.name AND packages.repo = :packages.repo"},
		{"?name=bash*", "WHERE packages.name LIKE :packages.name"},
		{"?name=bas?", "WHERE packages.name LIKE :packages.name"},
		{"?name=bash*&repo=", "WHERE packages.name LIKE :packages.name"},
		{"?name=bash*&repo=&arch=", "WHERE packages.name LIKE :packages.name"},
		{"?name=bash*&repo=main", "WHERE packages.name LIKE :packages.name AND packages.repo = :packages.repo"},
		{"?name=bash*&repo=main&arch=x86", "WHERE packages.arch = :packages.arch AND packages.name LIKE :packages.name AND packages.repo = :packages.repo"},
	}

	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			u, err := url.Parse(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			p := packagesWhere(u.Query())
			got := p["clause"]
			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
