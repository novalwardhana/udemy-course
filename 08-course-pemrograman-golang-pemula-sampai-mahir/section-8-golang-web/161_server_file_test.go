package belajar_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"sync"
	"testing"
)

// serveFile:
func serveFile(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		http.ServeFile(writer, request, "./resources/failed.html")
	} else {
		http.ServeFile(writer, request, "./resources/success.html")
	}
}

// TestServeFile:
func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(serveFile),
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

//go:embed resources/success.html
var resourcesSuccess string

//go:embed resources/failed.html
var resourcesFailed string

// embedServeFile:
func TestEmbedServeFile(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		firstName := request.URL.Query().Get("firstName")
		lastName := request.URL.Query().Get("lastName")

		if firstName == "" || lastName == "" {
			writer.Write([]byte(resourcesFailed))
			return
		}
		writer.Write([]byte(resourcesSuccess))
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// TestMultipleServeFile:
func TestMultipleServeFile(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	/* HTTP 1 */
	go func() {
		defer wg.Done()
		fmt.Println("Test HTTP 1")
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			email := request.URL.Query().Get("email")
			if email == "" {
				fmt.Fprint(writer, resourcesFailed)
				return
			}
			fmt.Fprint(writer, resourcesSuccess)
		}

		server := http.Server{
			Addr:    "localhost:8080",
			Handler: handler,
		}
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	/* HTTP 2 */
	go func() {
		defer wg.Done()
		fmt.Println("Test HTTP 2")
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			email := request.URL.Query().Get("email")
			if email == "" {
				http.ServeFile(writer, request, "./resources/failed.html")
				return
			}
			http.ServeFile(writer, request, "./resources/success.html")
		}

		server := http.Server{
			Addr:    "localhost:8081",
			Handler: handler,
		}
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	fmt.Println("Finish")
}
