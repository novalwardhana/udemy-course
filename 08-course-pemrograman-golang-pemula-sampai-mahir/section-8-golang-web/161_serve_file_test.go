package belajar_golang_web

import (
	"embed"
	"fmt"
	"net/http"
	"testing"
)

//go:embed resources/index.css
var resourcesIndex embed.FS

//go:embed resources/valid.html
var valid string

//go:embed resources/invalid.html
var invalid string

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("valid") == "yes" {
		http.ServeFile(writer, request, "./resources/valid.html")
	} else {
		http.ServeFile(writer, request, "./resources/invalid.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}
	server.ListenAndServe()
}

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("valid") == "yes" {
		fmt.Fprintf(writer, valid)
	} else {
		fmt.Fprintf(writer, invalid)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

//go:embed resources/mu.html
var mu string

//go:embed resources/chelsea.html
var chelsea string

//go:embed resources/arsenal.html
var arsenal string

//go:embed resources/noclub.html
var noclub string

func ServeFileEmbedAdvance(writer http.ResponseWriter, request *http.Request) {
	club := request.URL.Query().Get("club")
	switch club {
	case "mu":
		{
			fmt.Fprintf(writer, mu)
		}
	case "chelsea":
		{
			fmt.Fprintf(writer, chelsea)
		}
	case "arsenal":
		{
			fmt.Fprintf(writer, arsenal)
		}
	default:
		{
			fmt.Fprintf(writer, noclub)
		}
	}
}

func TestServeFileEmbedAdvance(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbedAdvance),
	}
	server.ListenAndServe()
}
