package belajar_golang_web

import (
	"fmt"
	"net/http"
	"path/filepath"
	"sync"
	"testing"
)

// downloadFile:
func downloadFile(writer http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")

	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "Bad request")
		return
	}

	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(writer, request, filepath.Join("./resources", file))
}

// TestDownloadFile:
func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(downloadFile),
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// TestMultipleDownloadFile:
func TestMultipleDownloadFile(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(3)

	/* Without attachment */
	go func() {
		defer wg.Done()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			file := request.URL.Query().Get("file")
			if file == "" {
				writer.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(writer, "Bad request")
				return
			}
			http.ServeFile(writer, request, filepath.Join("./resources", file))
		}
		server := http.Server{
			Addr:    "localhost:8080",
			Handler: handler,
		}

		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	/* With attachment */
	go func() {
		defer wg.Done()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			file := request.URL.Query().Get("file")
			if file == "" {
				writer.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(writer, "Bad request")
				return
			}
			writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file))
			http.ServeFile(writer, request, filepath.Join("./resources", file))
		}
		server := http.Server{
			Addr:    "localhost:8081",
			Handler: handler,
		}

		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	/* With attachment */
	go func() {
		defer wg.Done()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			file := request.URL.Query().Get("file")
			if file == "" {
				writer.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(writer, "Bad Request")
				return
			}
			writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file))
			http.ServeFile(writer, request, filepath.Join("./templates", file))
		}
		server := http.Server{
			Addr:    "localhost:8082",
			Handler: handler,
		}
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	fmt.Println("Finish")
}
