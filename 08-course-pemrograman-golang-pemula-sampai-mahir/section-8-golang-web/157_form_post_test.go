package belajar_golang_web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// formPost:
func formPost(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		panic(err)
	}
	firstName := request.PostForm.Get("firstName")
	lastName := request.PostForm.Get("lastName")
	fmt.Fprintln(writer, "first name: ", firstName)
	fmt.Fprintln(writer, "last name: ", lastName)
}

// TestFormPost:
func TestFormPost(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader("firstName=Noval&lastName=Wardhana"))
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	formPost(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// formPostMultipleValue:
func formPostMultipleValue(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		panic(err)
	}

	foods := request.PostForm["food"]
	beverages := request.PostForm["beverage"]
	breads := request.PostForm["bread"]
	fmt.Fprintln(writer, "Foods: ", strings.Join(foods, ", "))
	fmt.Fprintln(writer, "beverages: ", strings.Join(beverages, ", "))
	fmt.Fprintln(writer, "breads: ", strings.Join(breads, ", "))
}

// TestFormPostMultipleValue:
func TestFormPostMultipleValue(t *testing.T) {
	requestBody := strings.NewReader("bread=Hot%20Dog&food=Sate&food=Soto&food=NasiGoreng&food=Burger&beverage=CocaCola&beverage=Sprite&beverage=Teh")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	formPostMultipleValue(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// TestFormPostJSON:
func TestFormPostJSONString(t *testing.T) {
	type User struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Address string `json:"address"`
	}
	var users []User = []User{
		{"Noval", 29, "Bantul"},
		{"Novan", 30, "Kalimantan Timur"},
		{"Adrie", 31, "Kalimantan Selatan"},
		{"Deni", 30, "Jakarta Barat"},
	}
	byteJSON, _ := json.Marshal(users)
	requestBody := strings.NewReader(fmt.Sprintf("users=%s", string(byteJSON)))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	func(writer http.ResponseWriter, request *http.Request) {
		users := request.PostFormValue("users")
		fmt.Fprintln(writer, "Users:", users)
	}(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// TestFormPostJSON:
func TestFormPostJSONBody(t *testing.T) {
	type User struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Address string `json:"address"`
	}
	var users []User = []User{
		{"Noval", 29, "Bantul"},
		{"Novan", 30, "Kalimantan Timur"},
		{"Adrie", 31, "Kalimantan Selatan"},
		{"Deni", 30, "Jakarta Barat"},
	}
	byteUsers, _ := json.Marshal(users)
	requestBody := io.NopCloser(bytes.NewBuffer(byteUsers))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	recorder := httptest.NewRecorder()

	func(writer http.ResponseWriter, request *http.Request) {
		body := request.Body
		byteBody, err := io.ReadAll(body)
		if err != nil {
			fmt.Println("error 2")
			panic(err)
		}
		fmt.Fprintln(writer, "Users JSON Body: ", string(byteBody))
	}(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
