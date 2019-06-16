package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT
func users() http.HandlerFunc { // HLxxx
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "Handle GET") // HLxxx
		case http.MethodDelete:
			fmt.Fprintf(w, "Handle DELETE") // HLxxx
		case http.MethodPatch:
			fmt.Fprintf(w, "Handle PATCH") // HLxxx
		case http.MethodPost:
			fmt.Fprintf(w, "Handle POST") // HLxxx
		case http.MethodPut:
			fmt.Fprintf(w, "Handle PUT") // HLxxx
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

// END OMIT

func main() {
	http.HandleFunc("/users", users())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
