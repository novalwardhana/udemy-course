package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello world")
	}

	server := http.Server{
		Addr:    "localhost:9000",
		Handler: handler,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestHandlerAdvance(t *testing.T) {
	forever := make(chan int)
	go func() {
		go func() {
			var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("Novalita"))
			}
			server := http.Server{
				Addr:    "localhost:9001",
				Handler: handler,
			}
			if err := server.ListenAndServe(); err != nil {
				panic(err)
			}
		}()

		go func() {
			var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("Kusuma"))
			}
			server := http.Server{
				Addr:    "localhost:9002",
				Handler: handler,
			}
			if err := server.ListenAndServe(); err != nil {
				panic(err)
			}
		}()

		go func() {
			var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("Wardhana"))
			}
			server := http.Server{
				Addr:    "localhost:9003",
				Handler: handler,
			}
			if err := server.ListenAndServe(); err != nil {
				panic(err)
			}
		}()
	}()
	<-forever
}
