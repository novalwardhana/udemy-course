package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

// queryParameter:
func queryParameter(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	address := request.URL.Query().Get("address")
	if name != "" && address != "" {
		message := fmt.Sprintf("Hello %s from %s", name, address)
		writer.Write([]byte(message))
	} else {
		writer.Write([]byte("Hello"))
	}
}

// TestQueryParameter:
func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Noval&address=Bantul", nil)
	recorder := httptest.NewRecorder()

	queryParameter(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// multipleQueryParameter:
func multipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	name1 := request.URL.Query()["name1"]
	name2 := request.URL.Query()["name2"]
	name3 := request.URL.Query()["name3"]
	name4 := request.URL.Query()["name4"]
	name5 := request.URL.Query()["name5"]

	fmt.Fprintln(writer, "Name1:", strings.Join(name1, " "))
	fmt.Fprintln(writer, "Name2:", strings.Join(name2, " "))
	fmt.Fprintln(writer, "Name3:", strings.Join(name3, " "))
	fmt.Fprintln(writer, "Name4:", strings.Join(name4, " "))
	fmt.Fprintln(writer, "Name5:", strings.Join(name5, " "))
}

// TestMultipleQueryParameter:
func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name1=Noval&name2=Deni&name3=Adrie&name4=Novan&name1=Kusuma&name2=Eka&name3=Satrio&name4=Arief&name1=Wardhana&name2=Nugraha&name4=Kardianto", nil)
	recorder := httptest.NewRecorder()

	multipleQueryParameter(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// TestParallelQueryParameter:
func TestParallelQueryParameter(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(3)

	/* HTTP 1 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080?name=noval&age=29&address=Bantul", nil)
		recorder := httptest.NewRecorder()

		func(writer http.ResponseWriter, request *http.Request) {
			name := request.URL.Query().Get("name")
			age := request.URL.Query().Get("age")
			address := request.URL.Query().Get("address")
			fmt.Fprintln(writer, "HTTP 1 name: ", name)
			fmt.Fprintln(writer, "HTTP 1 age: ", age)
			fmt.Fprintln(writer, "HTTP 1 address: ", address)
		}(recorder, request)
		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	}()

	/* HTTP 2 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPut, "http://localhost:8080?club1=Arsenal&club1=Chelsea&club2=Liverpool&club2=RealMadrid&club2=Barcelona&club2=Sevilla&club3=Juventus&club3=Inter&club3=Milan", nil)
		recorder := httptest.NewRecorder()

		func(writer http.ResponseWriter, request *http.Request) {
			club1 := strings.Join(request.URL.Query()["club1"], ", ")
			club2 := strings.Join(request.URL.Query()["club2"], ", ")
			club3 := strings.Join(request.URL.Query()["club3"], ", ")
			fmt.Fprintln(writer, "HTTP 2 EPL club: ", club1)
			fmt.Fprintln(writer, "HTTP 2 LaLiga club: ", club2)
			fmt.Fprintln(writer, "HTTP 2 SeriA club: ", club3)
		}(recorder, request)
		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	}()

	/* HTTP 3 */
	go func() {
		defer wg.Done()

		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name1=noval&name1=Kusuma&name1=Wardhana", nil)
		recorder := httptest.NewRecorder()

		multipleQueryParameter(recorder, request)
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
