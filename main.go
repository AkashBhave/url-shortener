package main

import (
	"fmt"
	"net/http"
)

// Map of [alias] -> [URL]
var m = map[string]string{
	"go":    "https://golang.org",
	"jboss": "http://www.jboss.org/",
	"akash": "https://akashbhave.com",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(404)
			fmt.Fprintf(w, "alias not found")
			return
		}

		fmt.Fprintf(w, "welcome to Akash's URL shortener")
	})

	for alias, url := range m {
		http.HandleFunc("/"+alias, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, url, 301)
		})
	}

	http.ListenAndServe(":8800", nil)
}
