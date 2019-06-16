package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// START OMIT
func increment(count int) http.HandlerFunc { // HLxxx
	t := count + 1 // HLxxx
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method { // HLxxx
		case http.MethodGet:
			fmt.Fprintf(w, strconv.Itoa(t)) // HLxxx
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func main() {
	http.HandleFunc("/print2", increment(1))  // HLxxx
	http.HandleFunc("/print4", increment(3))  // HLxxx
	http.HandleFunc("/print10", increment(9)) // HLxxx

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END OMIT
