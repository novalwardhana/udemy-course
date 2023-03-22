package belajar_golang_web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// responseCode:
func responseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, "Name is empty")
	} else {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintln(writer, fmt.Sprintf("Hello %s", name))
	}
}

// TestResponseCode:
func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Noval", nil)
	recorder := httptest.NewRecorder()

	responseCode(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status: ", response.Status)
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Body: ", string(body))
}

// TestUserLoginResponseCode:
func TestUserLoginResponseCode(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(4)

	/* Login function */
	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var loginProcess = func(writer http.ResponseWriter, request *http.Request) {
		email := "noval@gmail.com"
		password := "test123"

		body := request.Body
		byteBody, err := io.ReadAll(body)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(writer, "Error parse parameters")
			return
		}
		var user User
		if err := json.Unmarshal(byteBody, &user); err != nil {
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(writer, "Error parse parameters")
			return
		}
		if user.Email == "" || user.Password == "" {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(writer, "Email or password is empty")
			return
		}
		if user.Email != email || user.Password != password {
			writer.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(writer, "Email or password not match")
			return
		}
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintln(writer, fmt.Sprintf("Success login | email: %s, password: %s", user.Email, user.Password))
	}

	/* HTTP 1 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()
		loginProcess(recorder, request)

		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("HTTP 1 status code: ", response.StatusCode)
		fmt.Println("HTTP 1 status: ", response.Status)
		fmt.Println("HTTP 1 message: ", string(body))
	}()

	/* HTTP 2 */
	go func() {
		defer wg.Done()
		var user = User{}
		userByte, _ := json.Marshal(user)
		requestBody := io.NopCloser(bytes.NewBuffer(userByte))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
		request.Header.Add("content-type", "application/json")
		recorder := httptest.NewRecorder()

		loginProcess(recorder, request)
		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("HTTP 2 status code: ", response.StatusCode)
		fmt.Println("HTTP 2 status: ", response.Status)
		fmt.Println("HTTP 2 message: ", string(body))
	}()

	/* HTTP 3 */
	go func() {
		defer wg.Done()
		var user = User{
			Email:    "noval@gmail.com",
			Password: "Haloo",
		}
		userByte, _ := json.Marshal(user)
		requestBody := io.NopCloser(bytes.NewBuffer(userByte))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
		request.Header.Add("content-type", "application/json")
		recorder := httptest.NewRecorder()

		loginProcess(recorder, request)
		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("HTTP 3 status code: ", response.StatusCode)
		fmt.Println("HTTP 3 status: ", response.Status)
		fmt.Println("HTTP 3 message: ", string(body))
	}()

	/* HTTP 4 */
	go func() {
		defer wg.Done()
		var user = User{
			Email:    "noval@gmail.com",
			Password: "test123",
		}
		byteUser, _ := json.Marshal(user)
		requestBody := io.NopCloser(bytes.NewBuffer(byteUser))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
		request.Header.Add("content-type", "application/json")
		recorder := httptest.NewRecorder()

		loginProcess(recorder, request)
		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("HTTP 4 status code: ", response.StatusCode)
		fmt.Println("HTTP 4 status: ", response.Status)
		fmt.Println("HTTP 4 message: ", string(body))
	}()

	wg.Wait()
	fmt.Println("Finish")
}
