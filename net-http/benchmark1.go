package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		output, _ := json.Marshal(map[string]string{"message": "pong"})
		w.Write(output)
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
