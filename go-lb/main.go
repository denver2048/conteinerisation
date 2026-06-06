package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello from %s 🚀\n", hostname)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}