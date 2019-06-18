package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleGreeting(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, format, "World")
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

// START OMIT
func adminOnly(h http.HandlerFunc) http.HandlerFunc { // HLxxx
	return func(w http.ResponseWriter, r *http.Request) {
		if 1 < 2 { // Auth Logic // HLxxx
			http.NotFound(w, r)
			return
		}
		h(w, r) // HLxxx
	}
}

func main() {
	http.HandleFunc("/formal", adminOnly(handleGreeting("Hello %s!"))) // HLxxx

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END OMIT
