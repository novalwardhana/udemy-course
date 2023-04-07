package main

import (
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

func TestMethodNotAllowedHandler(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Method not allowed")
	})
	router.POST("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprintf(writer, "POST")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Method not allowed", string(body))
}

// TestRedirectWhenMethodNotAllowed:
func TestRedirectWhenMethodNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.GET("/get-data/:email/:name/:web", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		email := params.ByName("email")
		name := params.ByName("name")
		web := params.ByName("web")
		message := fmt.Sprintf("email: %s | name: %s | web: %s", email, name, web)
		fmt.Fprintf(writer, message)
	})
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Route method not allowed")
	})

	wg := sync.WaitGroup{}
	wg.Add(2)

	/* Method allowed */
	go func() {
		defer wg.Done()
		requestMethod := http.MethodGet
		request := httptest.NewRequest(requestMethod, "http://localhost:8080/get-data/noval@gmail.com/novalwardhana/novalwardhana.com", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		assert.NotEqual(t, "Route method not allowed", string(body))
	}()

	/* Method not allowed */
	go func() {
		defer wg.Done()
		requestMethod := http.MethodPost
		request := httptest.NewRequest(requestMethod, "http://localhost:8080/get-data/noval@gmail.com/novalwardhana/novalwardhana.com", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		assert.Equal(t, "Route method not allowed", string(body))
	}()

	wg.Wait()
	fmt.Println("Finish")
}

// TestServerRequestNotAllowed:
func TestRedirectWhenRequestNotAllowed(t *testing.T) {
	fileSystem, _ := fs.Sub(resources, "resources")

	router := httprouter.New()
	router.POST("/get-data/:email/:name/:web", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		email := params.ByName("email")
		name := params.ByName("name")
		web := params.ByName("web")
		message := fmt.Sprintf("email: %s | name: %s | web: %s", email, name, web)
		fmt.Fprintf(writer, message)
	})
	router.GET("/redirect", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "general all page")
	})
	router.ServeFiles("/download/*filepath", http.FS(fileSystem))
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/redirect", http.StatusTemporaryRedirect)
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
