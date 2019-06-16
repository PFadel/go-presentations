package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT
func handleGreeting(format string) http.HandlerFunc { // HLxxx
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method { // HLxxx
		case http.MethodGet:
			fmt.Fprintf(w, format, "World") // HLxxx
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func main() {
	http.HandleFunc("/formal", handleGreeting("Hello %s!")) // HLxxx
	http.HandleFunc("/informal", handleGreeting("Hi %s!"))  // HLxxx
	http.HandleFunc("/sup", handleGreeting("'Sup %s!"))     // HLxxx

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END OMIT
