package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"

	"github.com/tosone/logging"

	_ "github.com/tosone/panx/statik"
)

func main() {
	var err error

	var f http.FileSystem
	if f, err = fs.New(); err != nil {
		log.Fatal(err)
	}

	var router = mux.NewRouter().StrictSlash(true)

	router.
		PathPrefix("/panx/").
		Handler(http.StripPrefix("/panx/", http.FileServer(f)))

	if err = http.ListenAndServe(":8080", router); err != nil {
		logging.Fatal("ListenAndServe Error: ", err)
	}
}
