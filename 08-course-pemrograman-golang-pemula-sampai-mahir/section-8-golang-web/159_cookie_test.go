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

// SetCookie:
func setCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "test-my-cookie"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprintln(writer, "Success add cookie")
}

// getCookie:
func getCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("test-my-cookie")
	if err != nil {
		fmt.Fprintln(writer, fmt.Sprintf("Error get cookie: %s", err.Error()))
		return
	}
	fmt.Fprintln(writer, fmt.Sprintf("Hello cookie: %s", cookie.Value))
}

// TestCookie:
func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", setCookie)
	mux.HandleFunc("/get-cookie", getCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// TestSetCookie:
func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=novalwardhana", nil)
	recorder := httptest.NewRecorder()

	setCookie(recorder, request)
	response := recorder.Result()
	for _, cookie := range response.Cookies() {
		message := fmt.Sprintf("Cookie %s: %s", cookie.Name, cookie.Value)
		fmt.Println(message)
	}
}

// TestGetCookie:
func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	cookie := new(http.Cookie)
	cookie.Name = "test-my-cookie"
	cookie.Value = "Noval"
	cookie.Path = "/"
	request.AddCookie(cookie)

	getCookie(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// TestBrowserMultipleCookie
func TestBrowserMultipleCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", func(writer http.ResponseWriter, request *http.Request) {
		cookie1 := new(http.Cookie)
		cookie1.Name = "username"
		cookie1.Value = "novalwardhana"
		cookie1.Path = "/"
		http.SetCookie(writer, cookie1)

		cookie2 := new(http.Cookie)
		cookie2.Name = "email"
		cookie2.Value = "noval@gmail.com"
		cookie2.Path = "/"
		http.SetCookie(writer, cookie2)

		cookie3 := new(http.Cookie)
		cookie3.Name = "name"
		cookie3.Value = "Novalita Kusuma Wardhana"
		cookie3.Path = "/"
		http.SetCookie(writer, cookie3)

		fmt.Fprintln(writer, "Success add cookie")
	})
	mux.HandleFunc("/get-cookie", func(writer http.ResponseWriter, request *http.Request) {
		for _, cookie := range request.Cookies() {
			fmt.Fprintln(writer, fmt.Sprintf("%s: %s", cookie.Name, cookie.Value))
		}
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// TestMultipleCookie:
func TestMultipleCookie(t *testing.T) {

	type User struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	/* Function set cookie */
	var setCookie = func(writer http.ResponseWriter, request *http.Request) {
		body := request.Body
		byteBody, err := io.ReadAll(body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(writer, err.Error())
			return
		}

		var user = new(User)
		if err := json.Unmarshal(byteBody, user); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(writer)
			return
		}

		emailCookie := new(http.Cookie)
		emailCookie.Name = "email"
		emailCookie.Value = user.Email
		emailCookie.Path = "/"
		http.SetCookie(writer, emailCookie)

		nameCookie := new(http.Cookie)
		nameCookie.Name = "name"
		nameCookie.Value = user.Name
		nameCookie.Path = "/"
		http.SetCookie(writer, nameCookie)

		passCookie := new(http.Cookie)
		passCookie.Name = "password"
		passCookie.Value = user.Password
		passCookie.Path = "/"
		http.SetCookie(writer, passCookie)

		fmt.Fprintf(writer, "HTTP 1 Success set cookie")
	}

	/* Function get cookie */
	var getCookie = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "HTTP 2 success get cookie")
		for _, cookie := range request.Cookies() {
			fmt.Fprintln(writer, fmt.Sprintf("%s: %s", cookie.Name, cookie.Value))
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	/* HTTP 1 set cookie */
	go func() {
		defer wg.Done()
		user := User{
			Email:    "noval@gmail.com",
			Name:     "Novalwardhana",
			Password: "test123",
		}
		byteUser, _ := json.Marshal(user)
		requestBody := io.NopCloser(bytes.NewBuffer(byteUser))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
		recorder := httptest.NewRecorder()

		setCookie(recorder, request)
		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
		for _, cookie := range response.Cookies() {
			cookieMessage := fmt.Sprintf("%s: %s", cookie.Name, cookie.Value)
			fmt.Println(cookieMessage)
		}
		fmt.Println("---")
	}()

	/* HTTP 2 get cookie */
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
		nameCookie := new(http.Cookie)
		nameCookie.Name = "name"
		nameCookie.Value = "Novalita Kusuma Wardhana"
		nameCookie.Path = "/"
		request.AddCookie(nameCookie)
		addressCookie := new(http.Cookie)
		addressCookie.Name = "address"
		addressCookie.Value = "Banguntapan"
		addressCookie.Path = "/"
		request.AddCookie(addressCookie)
		recorder := httptest.NewRecorder()

		getCookie(recorder, request)
		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	}()

	wg.Wait()
	fmt.Println("Finish")

}
