package models

import (
	"log"
	"net/url"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/myhro/aports-api/common"
	// Database Drivers are imported as blank imports
	_ "github.com/lib/pq"
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

	results := 50
	return results * (page - 1)
}
