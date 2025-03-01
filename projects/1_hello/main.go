package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, web!")
	})

	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
