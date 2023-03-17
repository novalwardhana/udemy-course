package belajar_golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestFileServer(t *testing.T) {
	dir := http.Dir("./resources")
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(dir)))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func TestFileServerEmbed(t *testing.T) {
	fileServer := http.FileServer(http.FS(resources))
	mux := http.NewServeMux()
	mux.Handle("/resources/", http.StripPrefix("/resources", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestFileServerAdvance(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.Handle("/test/", http.StripPrefix("/test", fileServer))
	mux.Handle("/trial/", http.StripPrefix("/trial", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
