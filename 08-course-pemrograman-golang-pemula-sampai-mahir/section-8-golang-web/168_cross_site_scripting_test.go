package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// templateAutoEscape:
func templateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "template_xss_simple.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>Ini adalah body <script>alert('Anda di heck')</script> </p>",
	})
}

// TestTemplateAutoEscape:
func TestTemplateAutoEscapeEnabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateAutoEscape(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// TestBrowserTemplateAutoEscape:
func TestTemplateAutoEscapeEnabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(templateAutoEscape),
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// templateAutoEscapeDisbaled:
func templateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "template_xss_simple.gohtml", map[string]interface{}{
		"Title": "Noval Web",
		"Body":  template.HTML("<h3>Welcome to my web</h3>"),
	})
}

// TestTemplateAutoEscapeDisabled:
func TestTemplateAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateAutoEscapeDisabled(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// TestTemplateAutoEscapeDisabledServer:
func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(templateAutoEscapeDisabled),
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// templateXSSExample:
func templateXSSAttack(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "template_xss_simple.gohtml", map[string]interface{}{
		"Title": "Noval Wab Page",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}

// TestTemplateXSSAttack:
func TestTemplateXSSAttack(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()
	templateXSSAttack(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// TestTemplateXSSAttackServer:
func TestTemplateXSSAttackServer(t *testing.T) {
	// http://localhost:8080/?body=<h1>Heck</h1><script>alert('anda di heck')</script>
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(templateXSSAttack),
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

//go:embed templates/template_xss_css_js_html.gohtml
var templateXSS embed.FS
var templateXSSLoader = template.Must(template.ParseFS(templateXSS, "templates/template_xss_css_js_html.gohtml"))

// TestMultipleTemplateXSS:
func TestMultipleTemplateXSS(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	/* Auto Escape Enable */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			templateXSSLoader.ExecuteTemplate(writer, "templateXSSAllSyntax", map[string]interface{}{
				"Title":  "NovalWardhana",
				"Css":    ".button{margin-top:120px}.navbar{top:10px;right:300px}",
				"Script": "<script>alert('Welcome to my blog')</script>",
				"Body":   "<p>Tampilan website noval</p>",
			})
		}
		handler(recorder, request)

		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	}()

	/* Auto Escape disable */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			templateXSSLoader.ExecuteTemplate(writer, "templateXSSAllSyntax", map[string]interface{}{
				"Title":  "Kusuma Wardhana",
				"Css":    template.CSS(".button{margin-top: 15px; background: #f2f2f2}.navbar{top: 10px; padding: 90px}"),
				"Script": template.HTML("<script>alert('Haloo bro...')</script>"),
				"Body":   template.HTML("<h2>Kusuma Wardhana</h2><p>Loren ipsum</p>"),
			})
		}
		handler(recorder, request)

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

// TestMultipleTemplateXSSServer:
func TestMultipleTemplateXSSServer(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	/*
		Auto escape disable
		http://localhost:8080/?body=<h1>Noval</h1><h2>Kusuma</h2><h3>Wardhana</h3><script>alert('Welcome to my web')</script>
	*/
	go func() {
		defer wg.Done()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			templateXSSLoader.ExecuteTemplate(writer, "templateXSSAllSyntax", map[string]interface{}{
				"Title":  "Noval Web 1.0",
				"Css":    template.CSS("<style type='text/css'>body{background: blue; padding: 100px;}</style>"),
				"Script": template.HTML("<script>alert('Welcome')</script>"),
				"Body":   template.HTML(request.URL.Query().Get("body")),
			})
		}

		server := http.Server{
			Addr:    "localhost:8080",
			Handler: handler,
		}
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	/*
		Auto escape enable
		http://localhost:8081/?body=<h1>Noval</h1><h2>Kusuma</h2><h3>Wardhana</h3><script>alert('Welcome to my web')</script>
	*/
	go func() {
		defer wg.Done()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			templateXSSLoader.ExecuteTemplate(writer, "templateXSSAllSyntax", map[string]interface{}{
				"Title":  "Noval Web 2.0",
				"Css":    "<style type='text/css'>body{background: blue; padding: 100px;}</style>",
				"Script": "<script>alert('Welcome')</script>",
				"Body":   request.URL.Query().Get("body"),
			})
		}

		server := http.Server{
			Addr:    "localhost:8081",
			Handler: handler,
		}
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	fmt.Println("Finish")
}
