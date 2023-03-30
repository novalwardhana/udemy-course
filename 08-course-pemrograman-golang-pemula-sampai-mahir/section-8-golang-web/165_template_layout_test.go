package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"text/template"
)

// templateLayout:
func templateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/template_layout_header.gohtml",
		"./templates/template_layout_body.gohtml",
		"./templates/template_layout_footer.gohtml",
	))
	t.ExecuteTemplate(writer, "body", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Noval wardhana",
	})
}

// TestTemplateLayout:
func TestTemplateLayout(t *testing.T) {
	fmt.Println("abc")
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateLayout(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func TestTemplateLayoutContainer(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(2)

	/* HTTP page 1 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			t := template.Must(template.ParseFiles("./templates/template_layout_design_1.gohtml", "./templates/template_layout_design_2.gohtml", "./templates/template_layout_design_3.gohtml"))
			t.ExecuteTemplate(writer, "template_layout_design_2.gohtml", map[string]interface{}{
				"Title": "Noval Web",
				"Menus": []string{"Home", "About", "Profile", "Contact", "Mail"},
				"Article": map[string]string{
					"Title": "Hello world",
					"Text":  "Loren ipsum dolor amet",
				},
				"Author": "novalwardhana",
				"Year":   "2023",
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

	/* HTTP Page 2 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			t := template.Must(template.ParseFiles("./templates/template_layout_design_1.gohtml", "./templates/template_layout_design_2.gohtml", "./templates/template_layout_design_3.gohtml"))
			t.ExecuteTemplate(writer, "body", map[string]interface{}{
				"Title": "Wardhana Blog",
				"Menus": []string{"Beranda", "Kontak", "Profil", "Email"},
				"Article": map[string]string{
					"Title": "Tutorial optimasi golang",
					"Text":  "Tutorial optimasi golang untuk linux",
				},
				"Author": "Wardhana",
				"Year":   "2022",
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
}
