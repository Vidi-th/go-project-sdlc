package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Dockerized Golang!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", nil)
}
