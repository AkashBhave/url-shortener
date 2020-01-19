package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome to Akash's URL shortener")
	})

	http.ListenAndServe(":8800", nil)
}
