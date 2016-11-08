package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

func main() {
	var HTTPAddr string
	flag.StringVar(&HTTPAddr, "http", "localhost:8080", "http address")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "OK")
	})

	log.Printf("listening on %s\n", HTTPAddr)
	log.Fatalln(http.ListenAndServe(HTTPAddr, nil))
}
