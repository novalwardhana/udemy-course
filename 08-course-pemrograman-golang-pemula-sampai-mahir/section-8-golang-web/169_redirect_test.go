package belajar_golang_web

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
	"testing"
)

// redirectFrom:
func redirectFrom(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Test redirect...!")
}

// redirectTo:
func redirectTo(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

// redirectOut:
func redirectOut(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "http://mi.com/id", http.StatusMovedPermanently)
}

// TestRedirect:
func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", redirectFrom)
	mux.HandleFunc("/redirect-to", redirectTo)
	mux.HandleFunc("/redirect-out", redirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// TestMultipleRedirect:
func TestMultipleRedirect(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	/* HTTP 1 */
	go func() {
		defer wg.Done()
		mux := http.NewServeMux()
		mux.HandleFunc("/redirect-from", func(writer http.ResponseWriter, request *http.Request) {
			http.Redirect(writer, request, "http://localhost:8081/redirect-to?title=Novalwardhana&body=<h1>Noval</h1><h2>Kusuma</h2></h3>Wardhana</h3>", http.StatusTemporaryRedirect)
		})
		mux.HandleFunc("/redirect-to", func(writer http.ResponseWriter, request *http.Request) {
			t := template.Must(template.ParseFiles("./templates/template_redirect.gohtml"))
			t.ExecuteTemplate(writer, "template_redirect.gohtml", map[string]interface{}{
				"Title": request.URL.Query()["title"][0],
				"Body":  template.HTML(request.URL.Query().Get("body")),
			})
		})
		mux.HandleFunc("/redirect-to-facebook", func(writer http.ResponseWriter, request *http.Request) {
			http.Redirect(writer, request, "http://facebook.com", http.StatusTemporaryRedirect)
		})

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
		mux := http.NewServeMux()
		mux.HandleFunc("/redirect-from", func(writer http.ResponseWriter, request *http.Request) {
			http.Redirect(writer, request, "http://localhost:8080/redirect-to?title=Halonoval&body=<script>alert('Redirect Example')</script>", http.StatusTemporaryRedirect)
		})
		mux.HandleFunc("/redirect-to", func(writer http.ResponseWriter, request *http.Request) {
			t := template.Must(template.ParseFiles("./templates/template_redirect.gohtml"))
			t.ExecuteTemplate(writer, "template_redirect.gohtml", map[string]interface{}{
				"Title": request.URL.Query().Get("title"),
				"Body":  template.HTML(request.URL.Query().Get("body")),
			})
		})

		server := http.Server{
			Addr:    "localhost:8081",
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	fmt.Println("Finish")
}
