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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Test Panic")
	})
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, i interface{}) {
		message := fmt.Sprintf("%v", i)
		fmt.Fprintf(writer, message)
	}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "Test Panic", string(body))

}

// TestMultiplePanicHandler:
func TestMultiplePanicHandler(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	/* Unit Test */
	go func() {
		defer wg.Done()
		type modelParams struct {
			Datas []float64
		}

		router := httprouter.New()
		router.POST("/data", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			var newParams modelParams
			body, _ := io.ReadAll(request.Body)
			json.Unmarshal(body, &newParams)

			var message string
			for i := 0; i < 10; i++ {
				message += fmt.Sprintf("Data index %d: %f\n", i, newParams.Datas[i])
			}
			fmt.Fprintf(writer, message)
		})
		router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, err interface{}) {
			message := fmt.Sprintf("Panic handler: %v", err)
			fmt.Fprintf(writer, message)
		}

		var params = modelParams{
			Datas: []float64{1, 4, 2, 9, 8, 10},
		}
		byteParams, _ := json.Marshal(params)
		requestBody := io.NopCloser(bytes.NewBuffer(byteParams))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/data", requestBody)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		assert.Equal(t, "Panic handler: runtime error: index out of range [6] with length 6", string(body))

	}()

	/* Server */
	go func() {
		defer wg.Done()
		router := httprouter.New()
		router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			fmt.Fprintf(writer, "Welcome to my API...")
		})
		router.POST("/average", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			body, err := io.ReadAll(request.Body)
			if err != nil {
				panic(err)
			}
			type modelParams struct {
				Datas []float64 `json:"datas"`
			}
			var newParams modelParams
			if err := json.Unmarshal(body, &newParams); err != nil {
				panic(err)
			}

			var sum float64
			var count float64
			for _, data := range newParams.Datas {
				sum += data
				count++
			}
			result := sum / count
			message := fmt.Sprintf("Params: %v \nResult: %f", newParams.Datas, result)
			fmt.Fprintf(writer, message)
		})
		router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, err interface{}) {
			message := fmt.Sprintf("An error occured: %v", err)
			fmt.Fprintf(writer, message)
		}

		server := http.Server{
			Addr:    "localhost:8080",
			Handler: router,
		}
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	fmt.Println("Finish")
}
