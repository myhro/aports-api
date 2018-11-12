package models

import (
	"fmt"
	"log"
	"net/url"
	"sort"
	"strings"
)

// Package represents a row in the packages table
type Package struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Description  string `json:"description"`
	URL          string `json:"url"`
	License      string `json:"license"`
	Repository   string `db:"repo" json:"repo"`
	Architecture string `db:"arch" json:"arch"`
	Maintainer   string `db:"maintainer_name" json:"maintainer"`
	BuildDate    string `db:"build_time" json:"build_date"`
}

// Packages returns an array of packages based on a filter
func Packages(params url.Values) []Package {
	db := connect()
	offset := getOffset(params)
	whereParams := packagesWhere(params)

	query := fmt.Sprintf(`
SELECT
	packages.name,
	packages.version,
	packages.description,
	packages.url,
	packages.license,
	packages.repo,
	packages.arch,
	maintainer.name AS maintainer_name,
	packages.build_time
FROM
	packages
INNER JOIN
	maintainer ON (packages.maintainer = maintainer.id)
%s
ORDER BY packages.build_time DESC
LIMIT 50
OFFSET %d`, whereParams["clause"], offset)

	pkgs := []Package{}
	rows, err := db.NamedQuery(query, whereParams)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		p := Package{}
		err := rows.StructScan(&p)
		if err != nil {
			log.Print(err)
		}
		pkgs = append(pkgs, p)
	}

	return pkgs
}

func packagesWhere(params url.Values) map[string]interface{} {
	defaultQuery := "WHERE "
	q := defaultQuery

	whereParams := map[string]interface{}{
		"packages.name":   params.Get("name"),
		"packages.repo":   params.Get("repo"),
		"packages.arch":   params.Get("arch"),
		"maintainer.name": params.Get("maintainer"),
	}

	keys := []string{}
	for k := range whereParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		if whereParams[k] == "" {
			delete(whereParams, k)
			continue
		}
		if strings.Contains(q, "=") || strings.Contains(q, "LIKE") {
			q += " AND "
		}
		if strings.ContainsAny(whereParams[k].(string), "*?") {
			whereParams[k] = replaceWildcards(whereParams[k].(string))
			q += fmt.Sprintf("%s LIKE :%s", k, k)
		} else {
			q += fmt.Sprintf("%s = :%s", k, k)
		}
	}

	whereParams["clause"] = ""
	if q != defaultQuery {
		whereParams["clause"] = q
	}

	return whereParams
}
