package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello world")
}

func TestHTTP(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	bodyByte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Result: ", string(bodyByte))
}

func TestHTTPNestedHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/test", nil)
	recorder := httptest.NewRecorder()

	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Welcome to my API"))
	}
	handler(recorder, request)

	response := recorder.Result()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status: ", response.Status)
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Result: ", string(responseBody))
}
