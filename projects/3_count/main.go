package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func main() {
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			mu.Lock()
			fmt.Fprintf(w, "Counter: %d", counter)
			mu.Unlock()
		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Invalid form data", http.StatusBadRequest)
				return
			}

			value := r.FormValue("count")
			num, err := strconv.Atoi(value)
			if err != nil {
				http.Error(w, "это не число", http.StatusBadRequest)
				return
			}

			mu.Lock()
			counter += num
			mu.Unlock()
			fmt.Fprintf(w, "Counter increased by %d", num)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Starting server on :3333...")
	http.ListenAndServe(":3333", nil)
}
