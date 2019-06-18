package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) { // HLxxx
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	s := &http.Server{
		Addr:         ":8080",          // HLxxx
		Handler:      nil,              // http.DefaultServeMux // HLxxx
		ReadTimeout:  10 * time.Second, // HLxxx
		WriteTimeout: 10 * time.Second, // HLxxx
	}

	log.Fatal(s.ListenAndServe()) // HLxxx
}
