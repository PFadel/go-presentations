package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// START OMIT
type Handler interface { // HLxxx
	ServeHTTP(http.ResponseWriter, *http.Request) // HLxxx
} // HLxxx

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.Handle("/foo", new(fooHandler)) // HLxxx

	log.Fatal(http.ListenAndServe(":8080", nil)) //  http.DefaultServeMux  // HLxxx
}

// END OMIT
