package belajar_golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Name is empty"))
	}
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "Name: %s", name)
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000?name=noval", nil)
	record := httptest.NewRecorder()
	ResponseCode(record, request)

	result := record.Result()
	resultBody, _ := ioutil.ReadAll(result.Body)
	fmt.Println("Result: ", string(resultBody))
	fmt.Println("Response Status: ", record.Result().Status)
	fmt.Println("Response Status Code: ", record.Result().StatusCode)

}

func ResponseCodeAdvance(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		panic(err)
	}
	firstName := request.PostForm.Get("first_name")
	midName := request.PostFormValue("mid_name")
	lastName := request.PostForm.Get("last_name")
	fmt.Println("firstName: ", firstName)
	if firstName == "" || midName == "" || lastName == "" {
		writer.WriteHeader(400)
		return
	}
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, firstName)
	fmt.Fprintln(writer, midName)
	fmt.Fprintln(writer, lastName)
}

func TestResponseCodeAdvance(t *testing.T) {
	requestBody := strings.NewReader("first_name=Novalita&mid_name=Kusuma&last_name=Wardhana")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000", requestBody)
	request.Header.Add("Content-type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()
	ResponseCodeAdvance(recorder, request)

	result := recorder.Result()
	resultBody, _ := ioutil.ReadAll(result.Body)
	fmt.Println(string(resultBody))
	fmt.Println(recorder.Result().StatusCode)
	fmt.Println(recorder.Result().Status)
}
