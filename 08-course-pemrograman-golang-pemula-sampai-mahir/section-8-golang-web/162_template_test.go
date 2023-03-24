package belajar_golang_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"text/template"
)

// simpleTemplateHTML1:
func simpleTemplateHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><head></head><body>{{.}}</body></html>`
	t, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "SIMPLE", "Test create template HTML 1")
}

// simpleTemplateHTML2:
func simpleTemplateHTML2(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><head></head><body>{{.}}</body></html>`
	t := template.Must(template.New("SIMPLE").Parse(templateText))
	t.ExecuteTemplate(writer, "SIMPLE", "Test create template HTML 2")
}

// TestTemplate:
func TestSimpleTemplateHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	simpleTemplateHTML2(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// fileTemplateHTML:
func fileTemplateHTML(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

// TestFileTemplateHTML:
func TestFileTemplateHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	fileTemplateHTML(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// fileTemplateDirectoryHTML:
func fileTemplateDirectoryHTML(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "simple.gohtml", "Test parse file gohtml")
}

// TestFileTemplateDirectoryHTML:
func TestFileTemplateDirectoryHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	fileTemplateDirectoryHTML(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var goHTMLTemplates embed.FS

// fileTemplateEmbed:
func fileTemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(goHTMLTemplates, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Test embed file template")
}

// TestFileTemplateEmbed:
func TestFileTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	fileTemplateEmbed(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// TestMultipleFileTemplate:
func TestMultipleFileTemplate(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(4)

	/* HTTP 1 */
	go func() {
		defer wg.Done()

		var processTemplate http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			templateHTML := `<html><head>{{.}}</head><body>{{.}}</body></html>`
			t := template.Must(template.New("simple.html").Parse(templateHTML))
			t.ExecuteTemplate(writer, "simple.html", "HTTP 1 template HTML")
		}

		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()
		processTemplate(recorder, request)

		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
		fmt.Println("")
	}()

	/* HTTP 2 */
	go func() {
		defer wg.Done()

		var processTemplate http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
			t.ExecuteTemplate(writer, "simple.gohtml", "HTTP 2 template HTML")
		}

		request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
		recorder := httptest.NewRecorder()
		processTemplate(recorder, request)

		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
		fmt.Println("")
	}()

	/* HTTP 3 */
	go func() {
		defer wg.Done()

		var processTemplate http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			t := template.Must(template.ParseGlob("./templates/*.gohtml"))
			t.ExecuteTemplate(writer, "simple.gohtml", "HTTP 3 template HTML")
		}

		request := httptest.NewRequest(http.MethodGet, "http://localhost:8082", nil)
		recorder := httptest.NewRecorder()
		processTemplate(recorder, request)

		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
		fmt.Println("")
	}()

	/* HTTP 4 */
	go func() {
		defer wg.Done()

		var processTemplate http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			t := template.Must(template.ParseFS(goHTMLTemplates, "templates/*.gohtml"))
			t.ExecuteTemplate(writer, "simple.gohtml", "HTTP 4 template HTML")
		}

		request := httptest.NewRequest(http.MethodGet, "http://localhost:8083", nil)
		recorder := httptest.NewRecorder()
		processTemplate(recorder, request)

		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
		fmt.Println("")
	}()

	wg.Wait()
	fmt.Println("Finish")
}
