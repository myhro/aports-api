package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/myhro/aports-api/common"
	"github.com/myhro/aports-api/packages"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Alpine Linux package database API")
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/packages", packages.ListHandler)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		access := fmt.Sprintf(`%s "%s %s" "%s"`, r.RemoteAddr, r.Method, r.URL.Path, r.UserAgent())
		log.Print(access)
		w.Header().Set("Content-Type", "application/json")
		router.ServeHTTP(w, r)
	})

	addr := common.GetEnv("API_ADDR", ":8000")
	srv := &http.Server{
		Handler: handler,
		Addr:    addr,
	}

	log.Print("Starting server on address ", addr)
	log.Fatal(srv.ListenAndServe())
}
