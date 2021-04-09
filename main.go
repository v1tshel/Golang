package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("Port")

	http.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "World")
	})

	http.ListenAndServe(":"+port, nil)
}
