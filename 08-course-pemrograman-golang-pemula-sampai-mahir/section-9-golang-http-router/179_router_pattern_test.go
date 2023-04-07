package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// TestRouterPatternNamedParameter:
func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/item/:itemID", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := fmt.Sprintf("Product %s Item %s", params.ByName("id"), params.ByName("itemID"))
		fmt.Fprintf(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/product/1/item/25", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 Item 25", string(body))
}

// TestRouterPatternCatchALlParameter:
func TestRouterPatternCatchALlParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/clubs/*clubParams", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := fmt.Sprintf("Clubs %s", params.ByName("clubParams"))
		fmt.Fprintf(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/clubs/arsenal/chelsea/juventus", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Clubs /arsenal/chelsea/juventus", string(body))
}

// TestCombineRouterPattern:
func TestCombineRouterPattern(t *testing.T) {
	router := httprouter.New()
	router.GET("/user/:name/*information", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		name := params.ByName("name")
		information := params.ByName("information")
		information = strings.Join(strings.Split(information, "/"), " ")
		message := fmt.Sprintf("User: %s, Information:%s", name, information)
		fmt.Fprintf(writer, message)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/user/novalwardhana/29/bantul/software-engineer", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "User: novalwardhana, Information: 29 bantul software-engineer", string(body))
}

// TestCombineRouterPatternServer:
func TestCombineRouterPatternServer(t *testing.T) {
	router := httprouter.New()
	router.GET("/user/:email/:name/*information", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		email := params.ByName("email")
		name := params.ByName("name")
		information := params.ByName("information")
		information = strings.Join(strings.Split(information, "/"), " ")
		message := fmt.Sprintf("email: %s, name: %s, information:%s", email, name, information)
		fmt.Fprintf(writer, message)
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
