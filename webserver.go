package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	startHTTPserver()
}

func startHTTPserver() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/list", listHandler)
	mux.HandleFunc("/details", detailsHandler)
	port := ":8090"
	fmt.Println("Starting, port is", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("Web server could not start, is another instance running?")
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index")
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Details")
}
