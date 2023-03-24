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

// templateData:
func templateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/template_data.gohtml"))
	t.ExecuteTemplate(writer, "template_data.gohtml", map[string]interface{}{
		"Title": "My Web",
		"Name":  "Noval Wardhana",
		"Email": "novalita.k.wardhana@gmail.com",
		"Address": map[string]interface{}{
			"City":       "Kota Yogyakarta",
			"PostalCode": 55290,
		},
	})
}

// TestTemplateData:
func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	templateDataMap(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// templateDataStruct:
func templateDataStruct(writer http.ResponseWriter, request *http.Request) {
	type Address struct {
		City       string
		PostalCode int
	}
	type User struct {
		Title   string
		Email   string
		Name    string
		Address Address
	}

	t := template.Must(template.ParseFiles("./templates/template_data.gohtml"))
	t.ExecuteTemplate(writer, "template_data.gohtml", User{
		Title: "User",
		Email: "noval@gmail.com",
		Name:  "Noval Wardhana",
		Address: Address{
			City:       "Bantul",
			PostalCode: 55198,
		},
	})
}

// TestTemplateDataStruct:
func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	templateDataStruct(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// TestTemplateDataBrowser:
func TestTemplateDataBrowser(t *testing.T) {
	type Address struct {
		City       string
		PostalCode int
	}
	type User struct {
		Title   string
		Email   string
		Name    string
		Address Address
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/chelsea", func(writer http.ResponseWriter, request *http.Request) {
		t := template.Must(template.ParseFiles("./templates/template_data.gohtml"))
		t.ExecuteTemplate(writer, "template_data.gohtml", User{
			Title: "Club",
			Email: "chelsea@gmail.com",
			Name:  "Chelsea FC",
			Address: Address{
				City:       "London",
				PostalCode: 98456,
			},
		})
	})
	mux.HandleFunc("/arsenal", func(writer http.ResponseWriter, request *http.Request) {
		t := template.Must(template.ParseFiles("./templates/template_data.gohtml"))
		t.ExecuteTemplate(writer, "template_data.gohtml", User{
			Title: "Club",
			Email: "arsenal@gmail.com",
			Name:  "Arsenal FC",
			Address: Address{
				City:       "North London",
				PostalCode: 98458,
			},
		})
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

//go:embed templates/template_data.gohtml
var templateData embed.FS

// TestMultipleTemplateData:
func TestMultipleTemplateData(t *testing.T) {

	type Address struct {
		City       string
		PostalCode int
	}
	type User struct {
		Title   string
		Email   string
		Name    string
		Address Address
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	/* HTTP 1 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()

		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			t := template.Must(template.ParseFS(templateData, "templates/template_data.gohtml"))
			t.ExecuteTemplate(writer, "template_data.gohtml", User{
				Title: "Club",
				Email: "manchester.united@gmail.com",
				Name:  "Manchester United",
				Address: Address{
					City:       "Manchester",
					PostalCode: 349867,
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
		fmt.Println("")
	}()

	/* HTTP 2 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8081", nil)
		recorder := httptest.NewRecorder()

		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			t := template.Must(template.ParseFS(templateData, "templates/template_data.gohtml"))
			t.ExecuteTemplate(writer, "template_data.gohtml", User{
				Title: "Club",
				Email: "liverpool@gmail.com",
				Name:  "Liverpool FC",
				Address: Address{
					City:       "Liverpool",
					PostalCode: 55198,
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
		fmt.Println("")
	}()

	wg.Wait()
	fmt.Println("Finish")
}
