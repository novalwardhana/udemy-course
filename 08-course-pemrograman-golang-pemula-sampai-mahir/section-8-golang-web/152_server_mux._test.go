package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("belajar server mux"))
	})
	mux.HandleFunc("/data/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test data noval")
	})
	mux.HandleFunc("/data/noval", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mas noval")
	})

	server := http.Server{
		Addr:    "localhost:9001",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestServerMuxAdvance(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome x"))
	})
	mux.HandleFunc("/team/one", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Arsenal"))
	})
	mux.HandleFunc("/team/two", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Chelsea"))
	})

	server := http.Server{
		Addr:    "localhost:9001",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
