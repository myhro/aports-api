package models

import (
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/myhro/aports-api/common"
	// Database Drivers are imported as blank imports
	_ "github.com/lib/pq"
)

const (
	// ResultsPerPage defines how many results will be shown by default
	ResultsPerPage = 50
)

func connect() *sqlx.DB {
	dbURL := common.GetEnv("API_DB_URL", "postgres://postgres@localhost/aports?sslmode=disable")
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getOffset(v url.Values) int {
	p := v.Get("page")
	if p == "" {
		return 0
	}

	page, err := strconv.Atoi(p)
	if err != nil {
		log.Print(err)
		return 0
	}
	if page < 1 {
		return 0
	}

	return ResultsPerPage * (page - 1)
}

func replaceWildcards(s string) string {
	s = strings.Replace(s, "*", "%", -1)
	s = strings.Replace(s, "?", "_", -1)
	return s
}
