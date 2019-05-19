// This is a simple HTTP server demo that uses net/http package only
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {

	port := ":8090"
	fmt.Println("Starting, port is", port)
	handler := http.HandlerFunc(routeHandler)
	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal("Web server could not start, is another instance running?")
	}
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")[1:]
	switch first := strings.ToLower(string(params[0])); first {
	case "list":
		fmt.Fprintf(w, "%s", first)
		listHandler(params[1:], w, r)
	case "details":
		fmt.Fprintf(w, "%s", first)
		detailsHandler(params[1:], w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("This page is not there"))
	}
}

func listHandler(params []string, w http.ResponseWriter, r *http.Request) {
	for k, v := range params {
		fmt.Printf("List param %d: %s\n", k, v)
	}
}

func detailsHandler(params []string, w http.ResponseWriter, r *http.Request) {
	for k, v := range params {
		fmt.Printf("Details param %d: %s\n", k, v)
	}
}
