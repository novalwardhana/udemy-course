package belajar_golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		panic(err)
	}
	name := request.PostForm.Get("name")
	age := request.PostForm.Get("age")
	address := request.PostForm.Get("address")
	fmt.Fprintf(writer, " form name: %s", name)
	fmt.Fprintf(writer, " form age: %s", age)
	fmt.Fprintf(writer, " form address: %s", address)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("name=Novalita Kusuma Wardhana&age=29")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()
	FormPost(recorder, request)

	result := recorder.Result()
	resultBody, _ := ioutil.ReadAll(result.Body)
	fmt.Println("Result: ", string(resultBody))
}

func FormPostAdvance(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		panic(err)
	}
	food := request.PostFormValue("food")
	price := request.PostFormValue("price")
	desc := request.PostFormValue("desc")

	fmt.Fprintf(writer, " food: %s \n", food)
	fmt.Fprintf(writer, " price: %s \n", price)
	fmt.Fprintf(writer, " desc: %s \n", desc)

	writer.Header().Add("Status", "Success")
	writer.Header().Add("Status Code", "200")
}

func TestFormPostAdvance(t *testing.T) {
	requestBody := strings.NewReader("food=Sate Ayam&price=15000")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:9000", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPostAdvance(recorder, request)
	result := recorder.Result()
	resultBody, _ := ioutil.ReadAll(result.Body)
	fmt.Println("Response: ", string(resultBody))

	fmt.Println("Response: ", recorder.Header().Get("Status"))
	fmt.Println("Response code: ", recorder.Header().Get("Status Code"))
}
