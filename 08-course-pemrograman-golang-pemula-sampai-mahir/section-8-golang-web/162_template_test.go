package belajar_golang_web

import (
	"embed"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func simpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	t := template.Must(template.New("simple").Parse(templateText))
	t.ExecuteTemplate(writer, "simple", "Test template success")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	simpleHTML(recorder, request)

	result := recorder.Result()
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println("Result: ", string(body))
}

func simpleHTMLAdvance(writer http.ResponseWriter, request *http.Request) {
	templateStr := `
		<html>
		<head>{{.}}</head>
		<body>
			<h1>{{.}}</h1>
			<p>{{.}}</p>
		</body>
		</html>
	`
	t := template.Must(template.New("sample.html").Parse(templateStr))
	t.ExecuteTemplate(writer, "sample.html", "My Name Noval Wardhana")
}

func TestSimpleHTMLAdvance(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	simpleHTMLAdvance(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response: ", string(body))
}

func templateHTML(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./resources/162_simple.gohtml"))
	t.ExecuteTemplate(writer, "162_simple.gohtml", "Test template success")
}

func TestTemplateHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateHTML(recorder, request)

	result := recorder.Result()
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println("Result: ", string(body))
}

func templateHTMLAdvance(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./resources/162_advance.gohtml"))
	t.ExecuteTemplate(writer, "162_advance.gohtml", "Novalita Kusuma Wardhana")
}

func TestTemplateHTMLAdvance(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateHTMLAdvance(recorder, request)

	result := recorder.Result()
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println("Result: ", string(body))
}

func templateDirectory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./resources/*.gohtml"))
	t.ExecuteTemplate(writer, "162_simple.gohtml", "Noval Wardhana")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateDirectory(recorder, request)

	result := recorder.Result()
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println("Result: ", string(body))
}

func templateDirectoryAdvance(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./resources/*.gohtml"))
	t.ExecuteTemplate(writer, "162_advance.gohtml", "Wardhana")
}

func TestTemplateDirectoryAdvance(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateDirectoryAdvance(recorder, request)

	result := recorder.Result()
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println("Result: ", string(body))
}

//go:embed resources/*.gohtml
var resourceTemplates embed.FS

func templateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(resourceTemplates, "resources/*.gohtml"))
	t.ExecuteTemplate(writer, "162_simple.gohtml", "Halonoval")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateEmbed(recorder, request)

	result := recorder.Result()
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println("Result: ", string(body))
}

//go:embed resources/*gohtml
var resourcesAdvanceTemplates embed.FS

func templateEmbedAdvance(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(resourcesAdvanceTemplates, "resources/*.gohtml"))
	t.ExecuteTemplate(writer, "162_advance.gohtml", "Mas noval")
}

func TestTemplateEmbedAdvance(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateEmbedAdvance(recorder, request)

	result := recorder.Result()
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println("Result: ", string(body))
}
