package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {

	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Hello World")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello World", string(body))
}

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// TestMultipleRouter:
func TestMultipleRouter(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	/* Router 1 */
	go func() {
		defer wg.Done()
		router := httprouter.New()
		router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			fmt.Fprintf(writer, fmt.Sprintf("Test http router 1 %s", request.URL.Query().Get("name")))
		})

		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Noval", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		fmt.Println("Test router 1")
		assert.Equal(t, "Test http router 1 Noval", string(body))
	}()

	/* Router 2 */
	go func() {
		defer wg.Done()

		router := httprouter.New()
		router.POST("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			body := request.Body
			byteBody, _ := io.ReadAll(body)
			user := new(User)
			json.Unmarshal(byteBody, user)
			fmt.Fprintf(writer, fmt.Sprintf("User email: %s, name: %s", user.Email, user.Name))
		})

		user := User{
			Email: "noval@gmail.com",
			Name:  "Noval",
		}
		userByte, _ := json.Marshal(&user)
		requestBody := io.NopCloser(bytes.NewBuffer(userByte))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8081/", requestBody)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		fmt.Println("Test router 2")
		assert.Equal(t, "User email: noval@gmail.com, name: Noval", string(body))
	}()

	wg.Wait()
	fmt.Println("Finish")
}
