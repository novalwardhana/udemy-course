package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Test not found")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/user", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "Test not found", string(body))
}

// TestMultipleNotFound:
func TestMultipleNotFound(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(2)

	/* Unit Test */
	go func() {
		defer wg.Done()
		router := httprouter.New()
		router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "Route not found")
		})

		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/user/test", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		assert.Equal(t, "Route not found", string(body))
	}()

	/* Server */
	go func() {
		defer wg.Done()

		router := httprouter.New()
		router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "Route not found")
		})

		server := http.Server{
			Addr:    "localhost:8080",
			Handler: router,
		}
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	/* Redirect */
	go func() {
		defer wg.Done()

		router := httprouter.New()
		router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			http.Redirect(writer, request, "http://mi.com/id", http.StatusTemporaryRedirect)
		})

		server := http.Server{
			Addr:    "localhost:8081",
			Handler: router,
		}
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	fmt.Println("Finish")
}
