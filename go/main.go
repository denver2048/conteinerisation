package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World from HERN1k with scratch docker image! 🚀\n"))
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server started on :8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}