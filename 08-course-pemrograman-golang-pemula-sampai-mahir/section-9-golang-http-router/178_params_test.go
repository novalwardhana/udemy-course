package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := fmt.Sprintf("product %s", params.ByName("id"))
		fmt.Fprintf(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/product/2", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "product 2", string(body))
}

// TestMultipleParams:
func TestMultipleParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/user/:email/:name/:address", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		email := params.ByName("email")
		name := params.ByName("name")
		address := params.ByName("address")
		message := fmt.Sprintf("User %s %s %s", email, name, address)
		fmt.Fprintf(writer, message)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/user/noval@gmail.com/novalwardhana/jogja", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "User noval@gmail.com novalwardhana jogja", string(body))
}
