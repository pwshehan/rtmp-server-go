package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		streamKey := r.URL.Query().Get("name")

		validStreamKeys := map[string]bool{
			"validstreamkey1": true,
			"validstreamkey2": true,
		}

		if validStreamKeys[streamKey] {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "OK")
		} else {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Invalid stream key")
		}
	})

	fmt.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
