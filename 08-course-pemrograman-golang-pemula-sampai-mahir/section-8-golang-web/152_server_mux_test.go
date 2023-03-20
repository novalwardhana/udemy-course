package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Main endpoint")
	})
	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Test endpoint")
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Images router")
	})
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Images thumbnails router")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestMultipleServeMux(t *testing.T) {
	forever := make(chan int)
	go func() {

		/* Server 1 */
		go func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
				fmt.Fprintf(writer, "Welcome...")
			})
			mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
				fmt.Fprintf(writer, "Test...")
			})

			server := http.Server{
				Addr:    "localhost:8080",
				Handler: mux,
			}
			if err := server.ListenAndServe(); err != nil {
				panic(err)
			}
		}()

		/* Server 2 */
		go func() {
			var mainHandler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("Main handler"))
			}
			var testHandler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("Test handler"))
			}
			mux := http.NewServeMux()
			mux.HandleFunc("/", mainHandler)
			mux.HandleFunc("/test", testHandler)

			server := http.Server{
				Addr:    "localhost:8081",
				Handler: mux,
			}
			if err := server.ListenAndServe(); err != nil {
				panic(err)
			}
		}()

		/* Server 3 */
		go func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
				name := request.URL.Query().Get("name")
				message := fmt.Sprintf("Hello %s", name)
				writer.Write([]byte(message))
			})

			server := http.Server{
				Addr:    "localhost:8082",
				Handler: mux,
			}
			if err := server.ListenAndServe(); err != nil {
				panic(err)
			}
		}()
	}()
	<-forever
}
