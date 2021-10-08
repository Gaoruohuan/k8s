package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

var goPath string = os.Getenv("GOROOT")

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Logging(index))
	mux.HandleFunc("/healthz", Logging(health))
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}

func index(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, strings.Join(v, ";"))
	}
	if goPath != "" {
		w.Header().Set("GOROOT", goPath)
	}
	w.WriteHeader(http.StatusOK)
}
