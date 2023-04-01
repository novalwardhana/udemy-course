package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS
var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func templateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Noval test")
}

// TestTemplateCaching:
func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateCaching(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templateCache embed.FS
var myTemplateCache = template.Must(template.ParseFS(templateCache, "templates/*.gohtml"))

type User struct {
	Title string
}

func (u User) ToUppercase() string {
	return strings.ToUpper(u.Title)
}

func (u User) ToLowercase() string {
	return strings.ToLower(u.Title)
}

func (u User) Message(hub string) string {
	arr := strings.Split(u.Title, "")
	return strings.Join(arr, hub)
}

// TestMultipleTemplateCaching:
func TestMultipleTemplateCaching(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(4)

	/* Template 1 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			myTemplateCache.ExecuteTemplate(writer, "simple.gohtml", "Noval Wardhana Home Page")
		}
		handler(recorder, request)

		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	}()

	/* Template 2 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			myTemplateCache.ExecuteTemplate(writer, "template_action_range.gohtml", map[string]interface{}{
				"Foods": []string{"Sate", "Bakso", "Mie Ayam", "Gulai", "Capcay", "Lumpia"},
				"Clubs": []string{"Arsenal", "Barcelona", "Liverpool", "Manchester United", "Real Madrid"},
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

	/* Template 3 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			myTemplateCache.ExecuteTemplate(writer, "template_action_with.gohtml", map[string]interface{}{
				"Name": "Novalita Kusuma Wardhana",
				"Address": map[string]string{
					"Street": "Jalan Wonocatur",
					"City":   "Bantul",
				},
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

	/* Template 4 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			myTemplateCache.ExecuteTemplate(writer, "templateFunction", User{
				Title: "NovalitaKusumaWardhana",
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
