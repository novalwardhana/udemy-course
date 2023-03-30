package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (m MyPage) SayHello(guest string) string {
	message := fmt.Sprintf("Hello %s, my name is %s", guest, m.Name)
	return message
}

// templateFunctionMyPage:
func templateFunctionMyPage(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Casca" }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{Name: "Noval"})
}

// TestTemplateFunctionMyPage:
func TestTemplateFunctionMyPage(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateFunctionMyPage(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// templateFunctionGLobal:
func templateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("Test").Parse(`{{len .Name}} {{len .Title}}`))
	t.ExecuteTemplate(writer, "Test", map[string]interface{}{
		"Title": "Haloo",
		"Name":  "Novalwardhana",
	})
}

// TestTemplateFunctionGLobal:
func TestTemplateFunctionGLobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateFunctionGlobal(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// templateFunctionCustom:
func templateFunctionCustom(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"ToUpper": func(param string) string {
			return strings.ToUpper(param)
		},
		"ToLower": func(param string) string {
			return strings.ToLower(param)
		},
	})
	t = template.Must(t.Parse(`{{ ToUpper .Name}} {{ToLower .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Noval kuSUMa WARdhaNa",
	})
}

// TestTemplateFunctionCustom:
func TestTemplateFunctionCustom(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateFunctionCustom(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// templateFunctionPipelines:
func templateFunctionPipelines(writer http.ResponseWriter, request *http.Request) {
	t := template.New("Pipelines")
	t = t.Funcs(map[string]interface{}{
		"greeting": func(param string) string {
			message := fmt.Sprintf("Hello bro %s", param)
			return message
		},
		"toUpper": func(param string) string {
			return strings.ToUpper(param)
		},
		"splitArray": func(param string) []string {
			return strings.Split(param, " ")
		},
	})
	t = template.Must(t.Parse(`{{ greeting .Name | toUpper | splitArray}}`))
	t.ExecuteTemplate(writer, "Pipelines", map[string]interface{}{
		"Name": "NovalWardhana",
	})
}

// TestTemplateFunctionPipelines:
func TestTemplateFunctionPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	templateFunctionPipelines(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

type TemplateFunction struct {
	Title string
}

func (t TemplateFunction) ToUppercase() string {
	return strings.ToUpper(t.Title)
}

func (t TemplateFunction) ToLowercase() string {
	return strings.ToLower(t.Title)
}

func (t TemplateFunction) Message(hub string) string {
	message := strings.ReplaceAll(t.Title, "", "-")
	return message
}

func TestTemplateFunctionGohtml(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(3)

	/* Template function 1 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			param := TemplateFunction{
				Title: "noVAlwarDHAana",
			}
			t := template.Must(template.ParseFiles("./templates/template_function_struct_method.gohtml"))
			t.ExecuteTemplate(writer, "templateFunction", param)
		}
		handler(recorder, request)

		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	}()

	/* Template function 2 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			t := template.Must(template.ParseFiles("./templates/template_function_global.gohtml"))
			t.ExecuteTemplate(writer, "templateFunction", map[string]interface{}{
				"Title":  "KusumaWardhana",
				"Name":   "Novalita Kusuma Wardhana",
				"Value1": 20,
				"Value2": 35,
				"Value3": 45,
				"Value4": 25,
				"Clubs":  []string{"Manchester United", "Barcelona", "Chelsea", "Arsenal"},
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

	/* Template function 3 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()
		var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
			t := template.New("templateFunction")
			t = t.Funcs(map[string]interface{}{
				"sumData": func(datas []float64) float64 {
					var sum float64
					for _, data := range datas {
						sum += data
					}
					return sum
				},
				"meanData": func(datas []float64) float64 {
					var sum float64
					var length float64
					for _, data := range datas {
						sum += data
						length++
					}
					return sum / length
				},
			})
			t = template.Must(t.Parse(`
				sum: {{sumData .Datas}}
				mean: {{meanData .Datas}}
			`))
			t.ExecuteTemplate(writer, "templateFunction", map[string]interface{}{
				"Datas": []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
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
