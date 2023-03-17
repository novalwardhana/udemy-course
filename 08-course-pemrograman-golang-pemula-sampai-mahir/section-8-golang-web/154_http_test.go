package belajar_golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HTTPHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello world")
}

func TestHTTP(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/test", nil)
	recorder := httptest.NewRecorder()

	HTTPHandler(recorder, request)
	response := recorder.Result()
	responseStr, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response; ", string(responseStr))
}

func HTTPHandlerAdvance(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Test handler")
	fmt.Fprintln(writer, "Test 2")
	fmt.Fprintln(writer, "Test 3")
}

func TestHTTPAdvance(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	HTTPHandlerAdvance(recorder, request)

	response := recorder.Result()
	responseStr, _ := ioutil.ReadAll(response.Body) // golang 1.16++ io.ReadAll()
	fmt.Println("Response; ", string(responseStr))
}
