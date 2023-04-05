package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execure handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After execute handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (eh *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Recover panic...")
		}
	}()
	eh.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler executed")
		fmt.Fprintf(writer, "Hello middleware")
	})
	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Test executed")
		fmt.Fprintf(writer, "Test route middleware")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		panic("Test")
	})
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

type ParamMiddleware struct {
	Handler http.Handler
}

func (p *ParamMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic handler")
			fmt.Fprintf(writer, fmt.Sprintf("Panic recover: %s", r))
		}
	}()
	email := request.Header.Get("email")
	password := request.Header.Get("password")
	if email == "" || password == "" {
		fmt.Fprintf(writer, "Missing header email and password")
		return
	}
	p.Handler.ServeHTTP(writer, request)
}

type AuthMiddleware struct {
	Handler http.Handler
}

func (a *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	email := request.Header.Get("email")
	password := request.Header.Get("password")
	if email != "noval@gmail.com" || password != "test123" {
		fmt.Fprintf(writer, "Unauthorized")
		return
	}
	a.Handler.ServeHTTP(writer, request)
}

func TestAuthMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Welcome to my API")
		fmt.Fprintf(writer, "Welcome to my API")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		panic("Test")
	})
	mux.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		file, fileHeader, err := request.FormFile("file")
		if err != nil {
			panic(err)
		}

		fileTarget, err := os.Create(filepath.Join("./resources", fileHeader.Filename))
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(fileTarget, file); err != nil {
			panic(err)
		}

		message := fmt.Sprintf("Hello %s, file successfully uploaded", request.FormValue("name"))
		fmt.Fprintf(writer, message)
	})
	mux.HandleFunc("/download", func(writer http.ResponseWriter, request *http.Request) {
		file := request.URL.Query().Get("file")
		if file == "" {
			panic("bad request")
		}

		writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file))
		http.ServeFile(writer, request, filepath.Join("./resources", file))
	})

	authMiddleware := new(AuthMiddleware)
	authMiddleware.Handler = mux

	paramMiddleware := new(ParamMiddleware)
	paramMiddleware.Handler = authMiddleware

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: paramMiddleware,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
