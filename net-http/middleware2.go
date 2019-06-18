package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type respWriter struct {
	http.ResponseWriter
	status int
	body   []byte
}

func (w *respWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *respWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	w.body = b
	return w.ResponseWriter.Write(b)
}

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
func logger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rec := respWriter{ResponseWriter: w} // HLxxx

		start := time.Now()
		h(&rec, r) // HLxxx
		elapsed := time.Since(start)
		log.Printf("Took %s with status code %d", elapsed.String(), rec.status) // HLxxx
	}
}

func main() {
	http.HandleFunc("/formal", logger(handleGreeting("Hello %s!"))) // HLxxx

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END OMIT
