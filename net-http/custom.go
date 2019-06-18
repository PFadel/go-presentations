package main

import (
	"net/http"
)

// START OMIT
type respWriter struct {
	http.ResponseWriter        // HLxxx
	status              int    // HLxxx
	body                []byte // HLxxx
}

func (w *respWriter) WriteHeader(status int) { // HLxxx
	w.status = status
	w.ResponseWriter.WriteHeader(status) // HLxxx
}

func (w *respWriter) Write(b []byte) (int, error) { // HLxxx
	if w.status == 0 {
		w.status = http.StatusOK
	}
	w.body = b
	return w.ResponseWriter.Write(b) // HLxxx
}

// END OMIT
