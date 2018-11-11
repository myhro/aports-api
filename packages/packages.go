package packages

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/myhro/aports-api/models"
)

func formatUnixTime(unix string) string {
	u, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	t := time.Unix(u, 0)
	return t.Format(time.RFC3339)
}

// ListHandler returns a list of packages
func ListHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	pkgs := models.Packages(params)
	for i := range pkgs {
		pkgs[i].BuildDate = formatUnixTime(pkgs[i].BuildDate)
	}
	res, err := json.Marshal(pkgs)
	if err != nil {
		log.Print(err)
	}
	fmt.Fprintf(w, string(res))
}
