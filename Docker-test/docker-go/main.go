package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Sanjana...")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8081", nil)
}

