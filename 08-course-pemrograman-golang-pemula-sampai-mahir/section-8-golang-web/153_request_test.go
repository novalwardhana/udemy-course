package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL)
		fmt.Fprintf(w, r.Method)
	}

	server := http.Server{
		Addr:    "localhost:9001",
		Handler: handler,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestRequestAdvance(t *testing.T) {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Method: ", r.Method)
		fmt.Fprintln(w, "URI: ", r.RequestURI)
		fmt.Fprintln(w, "URl: ", r.URL)
		fmt.Fprintln(w, "Host: ", r.Host)
	})

	server := http.Server{
		Addr:    "localhost:9002",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
