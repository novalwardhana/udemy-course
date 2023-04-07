package main

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type middleware struct {
	Handler http.Handler
}

func (m *middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Run Middleware 1")
	m.Handler.ServeHTTP(writer, request)
	fmt.Println("Run middleware 2")
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Welcome")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Welcome", string(body))
}

// BasicAuthMiddleware:
type BasicAuthMiddleware struct {
	Handler http.Handler
}

func (b *BasicAuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var correctUsername = "novalwardhana"
	var correctPassword = "test123"
	username, password, ok := request.BasicAuth()
	if !ok {
		fmt.Fprintf(writer, "Username or password no filled")
		return
	}
	if username != correctUsername || password != correctPassword {
		fmt.Fprintf(writer, "Username or password not match")
		return
	}

	b.Handler.ServeHTTP(writer, request)
}

// TestBasicAuthMiddleware:
func TestBasicAuthMiddleware(t *testing.T) {
	fileSystem, _ := fs.Sub(resources, "resources")

	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Welcome to my API")
	})
	router.POST("/upload", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		file, fileHeader, err := request.FormFile("file")
		if err != nil {
			panic(err)
		}
		fileDestination, err := os.Create(filepath.Join("./resources", fileHeader.Filename))
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(fileDestination, file); err != nil {
			panic(err)
		}

		fmt.Fprintf(writer, "Success upload file")
	})
	router.ServeFiles("/download/*filepath", http.FS(fileSystem))
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, err interface{}) {
		fmt.Fprintf(writer, fmt.Sprintf("Panic occured: %v", err))
	}
	basicAuthMiddleware := new(BasicAuthMiddleware)
	basicAuthMiddleware.Handler = router

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: basicAuthMiddleware,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
