package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "Name parameter is missing", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
