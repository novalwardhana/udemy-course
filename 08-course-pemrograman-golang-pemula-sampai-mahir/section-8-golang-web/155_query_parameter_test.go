package belajar_golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HTTPQueryParameter(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name != "" {
		fmt.Fprintf(writer, "Haloo %s", name)
	} else {
		fmt.Fprintf(writer, "Param empty")
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Noval", nil)
	recorder := httptest.NewRecorder()
	HTTPQueryParameter(recorder, request)

	response := recorder.Result()
	responseByte, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response: ", string(responseByte))
}

func HTTPMultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	age := request.URL.Query().Get("age")
	address := request.URL.Query().Get("address")
	fmt.Fprintf(writer, "Name: %s \n", name)
	fmt.Fprintf(writer, "Age: %s \n", age)
	fmt.Fprintf(writer, "Address: %s \n", address)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9001?name=noval&age=25&address=Bantul", nil)
	record := httptest.NewRecorder()
	HTTPMultipleQueryParameter(record, request)

	response := record.Body
	responseByte, _ := ioutil.ReadAll(response)
	fmt.Println("Response: ", string(responseByte))
}

func HTTPMultipleValueQueryParameter(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprintf(writer, "Response: %s", strings.Join(names, "-"))
}

func TestMultipleValueQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Noval&name=Kusuma&name=wardhana", nil)
	record := httptest.NewRecorder()
	HTTPMultipleValueQueryParameter(record, request)

	response := record.Body
	responseByte, _ := ioutil.ReadAll(response)
	fmt.Println("Response: ", string(responseByte))
}

func HTTPMultipleValueQueryParameterAdvance(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	tours := query["tour"]
	clubs := query["club"]
	smartphone := query["smartphone"]

	fmt.Fprintf(writer, "Tours: %s \n", strings.Join(tours, ","))
	fmt.Fprintf(writer, "Clubs: %s \n", strings.Join(clubs, ","))
	fmt.Fprintf(writer, "Smartphones: %s \n", strings.Join(smartphone, ","))
}

func TestMultipleValueQueryParameterAdvance(t *testing.T) {
	url := "http://localhost:8000?"
	url += "tour=gunungkidul&tour=kulonprogo&tour=bantul&"
	url += "club=mandhester-united&club=barcelona&club=manchester-city&"
	url += "smartphone=oppo&smartphone=vivo&smartphone=xiaomi"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	record := httptest.NewRecorder()
	HTTPMultipleValueQueryParameterAdvance(record, request)

	response := record.Body
	responseByte, _ := ioutil.ReadAll(response)
	fmt.Println("Response: ", string(responseByte))
}
