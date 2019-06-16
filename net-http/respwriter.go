package main

import "net/http"

// START OMIT
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
	n, err := w.ResponseWriter.Write(b)
	return n, err
}

// END OMIT
