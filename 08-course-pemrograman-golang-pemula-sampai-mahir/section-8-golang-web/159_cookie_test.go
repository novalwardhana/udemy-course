package belajar_golang_web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "myname"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprintf(writer, "success create cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("myname")
	if err != nil {
		fmt.Fprintf(writer, "No cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Haloo %s", name)
	}
}

func TestRunCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=novalwardhana", nil)
	recorder := httptest.NewRecorder()
	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Println(cookie.Name, cookie.Value, cookie.Path)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=ronaldo", nil)
	recorder := httptest.NewRecorder()
	cookie := new(http.Cookie)
	cookie.Name = "myname"
	cookie.Value = "Noval Wardhana"
	cookie.Path = "/"
	request.AddCookie(cookie)
	GetCookie(recorder, request)

	result := recorder.Result()
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println(string(body))
}

func SetCookieAdvance(writer http.ResponseWriter, request *http.Request) {
	type Params struct {
		FirstName string `json:"first_name"`
		MidName   string `json:"mid_name"`
		LastName  string `json:"last_name"`
	}
	var params Params
	body, _ := ioutil.ReadAll(request.Body)
	if err := json.Unmarshal(body, &params); err != nil {
		panic(err)
	}
	if params.FirstName != "" && params.MidName != "" && params.LastName != "" {
		cookie1 := new(http.Cookie)
		cookie1.Name = "cookie1"
		cookie1.Value = params.FirstName
		http.SetCookie(writer, cookie1)

		cookie2 := new(http.Cookie)
		cookie2.Name = "cookie2"
		cookie2.Value = params.MidName
		http.SetCookie(writer, cookie2)

		cookie3 := new(http.Cookie)
		cookie3.Name = "cookie3"
		cookie3.Value = params.LastName
		http.SetCookie(writer, cookie3)
		fmt.Fprintf(writer, "Success set cookie")
		return
	}
	fmt.Fprintf(writer, "Failed set cookie")
}

func GetCookieAdvance(writer http.ResponseWriter, request *http.Request) {
	cookie1, _ := request.Cookie("cookie1")
	fmt.Fprintln(writer, cookie1.Name, cookie1.Value)

	cookie2, _ := request.Cookie("cookie2")
	fmt.Fprintln(writer, cookie2.Name, cookie2.Value)

	cookie3, _ := request.Cookie("cookie3")
	fmt.Fprintln(writer, cookie3.Name, cookie3.Value)
}

func TestRunCookieAdvance(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookieAdvance)
	mux.HandleFunc("/get-cookie", GetCookieAdvance)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestSetCookieAdvance(t *testing.T) {
	type Params struct {
		FirstName string `json:"first_name"`
		MidName   string `json:"mid_name"`
		LastName  string `json:"last_name"`
	}
	var params Params
	params.FirstName = "Cascarilla"
	params.MidName = "Novi"
	params.LastName = "Wahyudha"
	byteParams, _ := json.Marshal(params)
	body := strings.NewReader(string(byteParams))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", body)
	recorder := httptest.NewRecorder()
	SetCookieAdvance(recorder, request)

	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Println(cookie.Name, cookie.Value)
	}
}

func TestGetCookieAdvance(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	cookie1 := new(http.Cookie)
	cookie1.Name = "cookie1"
	cookie1.Value = "Dhea"
	request.AddCookie(cookie1)

	cookie2 := new(http.Cookie)
	cookie2.Name = "cookie2"
	cookie2.Value = "Saintystya"
	request.AddCookie(cookie2)

	cookie3 := new(http.Cookie)
	cookie3.Name = "cookie3"
	cookie3.Value = "Brilliani"
	request.AddCookie(cookie3)

	GetCookieAdvance(recorder, request)

	result := recorder.Result()
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println(string(body))
}
