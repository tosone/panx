package main

import (
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"

	_ "github.com/tosone/panx/statik"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	http.ListenAndServe(":8080", nil)
}
