package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

// RequestHeader:
func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	timestamp := request.Header.Get("timestamp")
	fmt.Fprintln(writer, "Header request contentType: ", contentType)
	fmt.Fprintln(writer, "Header request timestamp: ", timestamp)
}

// TestHeader:
func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("content-type", "binary")
	request.Header.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// ResponseWriter:
func ResponseWriter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "NovalWardhana")
	fmt.Fprintln(writer, "Success")
}

// TestResponseWriter:
func TestResponseWriter(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseWriter(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	fmt.Println(recorder.Header().Get("x-powered-BY"))
}

// TestRequestResponseHeader:
func TestRequestResponseHeader(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	/* HTTP 1 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		request.Header.Add("username", "Noval")
		request.Header.Add("useRName", "Kusuma")
		request.Header.Add("userNAME", "Wardhana")
		request.Header.Add("content-type", "application")
		request.Header.Add("content-TYPE", "json")
		request.Header.Add("CONTENT-type", "binary")
		recorder := httptest.NewRecorder()

		func(writer http.ResponseWriter, request *http.Request) {
			username := request.Header.Values("username")
			contentType := request.Header.Values("content-type")
			fmt.Fprintln(writer, "HTTP 1 Header username: ", strings.Join(username, ", "))
			fmt.Fprintln(writer, "HTTP 1 Header content-type: ", strings.Join(contentType, ", "))
		}(recorder, request)
		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	}()

	/* HTTP 2 */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
		recorder := httptest.NewRecorder()

		func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
			writer.Header().Add("timeSTAMP", time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
			writer.Header().Add("TIMEstamp", time.Now().AddDate(0, 0, 1).Format("2006-01-02 15:04:05"))
			writer.Header().Add("Response-Type", "json")
			writer.Header().Add("responSE-TYpe", "protobuf")
			writer.Header().Add("response-type", "binary")
			writer.Write([]byte("HTTP 2 success add response header"))
		}(recorder, request)
		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
		timestamps := response.Header.Values("timestamp")
		responseType := response.Header.Values("response-type")
		fmt.Println("HTTP 2 response header timestamp: ", strings.Join(timestamps, ", "))
		fmt.Println("HTTP 2 response header response-type: ", strings.Join(responseType, ", "))
	}()

	wg.Wait()
	fmt.Println("Finish")
}
