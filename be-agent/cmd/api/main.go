package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting on :8080...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from be-agent!")
	})
	http.ListenAndServe(":8080", nil)
}
