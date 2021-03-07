package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// Writes the formatted string back to first argument
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}
func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8000", nil)
}
