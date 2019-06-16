package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// START OMIT
type HandlerFunc func(http.ResponseWriter, *http.Request) // HLxxx

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/foo", foo) // HLxxx

	log.Fatal(http.ListenAndServe(":8080", nil)) //  http.DefaultServeMux  // HLxxx
}

// END OMIT
