package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusBadGateway)
	w.WriteHeader(http.StatusServiceUnavailable)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
