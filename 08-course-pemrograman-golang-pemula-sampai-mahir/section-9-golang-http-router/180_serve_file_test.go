package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/images/*filepath", http.FS(directory))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/images/test.txt", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Noval Wardhana", string(body))
}

// TestMultipleServeFile:
func TestMultipleServeFile(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		fileSystem, err := fs.Sub(resources, "resources")
		if err != nil {
			panic(err)
		}
		mux := httprouter.New()
		mux.ServeFiles("/file/*filepath", http.FS(fileSystem))

		server := http.Server{
			Addr:    "localhost:8080",
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}

	}()

	go func() {
		defer wg.Done()
		fileSystem, err := fs.Sub(resources, "resources")
		if err != nil {
			panic(err)
		}
		router := httprouter.New()
		router.ServeFiles("/file/*filepath", http.FS(fileSystem))
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/file/real_madrid.png", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	}()

	wg.Wait()
	fmt.Println("Finish")
}
