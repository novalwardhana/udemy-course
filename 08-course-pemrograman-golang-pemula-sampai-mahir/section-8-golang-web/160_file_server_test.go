package belajar_golang_web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"sync"
	"testing"
)

// TestFileServer:
func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

// TestEmbedFileServer:
func TestEmbedFileServer(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// TestMultipleFileServer:
func TestMultipleFileServer(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	/* HTTP 1 */
	go func() {
		defer wg.Done()
		directory, _ := fs.Sub(resources, "resources")
		fileServer := http.FileServer(http.FS(directory))
		mux := http.NewServeMux()
		mux.Handle("/test/", http.StripPrefix("/test", fileServer))
		server := http.Server{
			Addr:    "localhost:8080",
			Handler: mux,
		}

		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	/* HTTP 2 */
	go func() {
		defer wg.Done()
		directory := http.Dir("./resources")
		fileServer := http.FileServer(directory)
		mux := http.NewServeMux()
		mux.Handle("/try/", http.StripPrefix("/try", fileServer))
		server := http.Server{
			Addr:    "localhost:8081",
			Handler: mux,
		}

		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	fmt.Println("Success")
}
