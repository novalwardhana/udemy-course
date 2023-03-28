package belajar_golang_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func templateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/template_action.gohtml"))
	t.ExecuteTemplate(writer, "template_action.gohtml", map[string]interface{}{
		"Name": "El lohim",
	})
}

//go:embed templates/template_action_if_else.gohtml
var templateActionIfElse embed.FS

func templateActionElseIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templateActionIfElse, "templates/template_action_if_else.gohtml"))
	t.ExecuteTemplate(writer, "template_action_if_else.gohtml", map[string]interface{}{
		"Title": "Haloo noval",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateActionIf(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func TestTemplateActionElseIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateActionElseIf(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func templateActionComparator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/template_action_comparator.gohtml"))
	t.ExecuteTemplate(writer, "template_action_comparator.gohtml", map[string]interface{}{
		"FinalValue": 91,
	})
}

func TestTemplateActionComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateActionComparator(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// templateActionRange:
func templateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/template_action_range.gohtml"))
	t.ExecuteTemplate(writer, "template_action_range.gohtml", map[string]interface{}{
		"Clubs": []string{"MU", "chelsea", "Arsenal"},
		"Foods": []string{"Soto", "Bakso", "Sate"},
	})
}

// TestTemplateActionRange:
func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateActionRange(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// templateActionWith:
func templateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/template_action_with.gohtml"))
	t.ExecuteTemplate(writer, "template_action_with.gohtml", map[string]interface{}{
		"Name": "Noval wardhana",
		"Address": map[string]interface{}{
			"Street": "Jalan Wonocatur",
			"City":   "Bantul",
		},
	})
}

// TestTemplateActionWith:
func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateActionWith(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// templateActionAllSyntax:
func templateActionAllSyntax(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/template_action_all_syntax.gohtml"))
	t.ExecuteTemplate(writer, "template_action_all_syntax.gohtml", map[string]interface{}{
		"Title":  "Template Action",
		"Name1":  "Noval",
		"Name2":  "Kusuma",
		"Name3":  "Wardhana",
		"Value1": 20,
		"Value2": 29,
		"clubs":  []string{"Real Madrid", "Manchester United", "Juventus"},
		"leagues": []interface{}{
			[]string{"AC Milan", "Inter Milan", "Napoli"},
			[]string{"Liverpool", "Leichester"},
		},
		"MyData": map[string]interface{}{
			"Email":    "noval@gmail.com",
			"username": "novalwardhana",
		},
	})
}

// TestTemplateActionAllSyntax:
func TestTemplateActionAllSyntax(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateActionAllSyntax(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
