package belajar_golang_web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello world")
	}
	var server = http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestMultipleHandler(t *testing.T) {
	forever := make(chan int)
	go func() {

		/* Server 1 */
		go func() {
			var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
				fmt.Fprintf(writer, "Novalita Kusuma Wardhana")
			}
			var server = http.Server{
				Addr:    "localhost:8080",
				Handler: handler,
			}
			if err := server.ListenAndServe(); err != nil {
				panic(err)
			}
		}()

		/* Server 2 */
		go func() {
			var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
				type User struct {
					Name    string `json:"name"`
					Age     int    `json:"age"`
					Address string `json:"address"`
				}
				var newUser User = User{
					Name:    "Noval Wardhana",
					Age:     29,
					Address: "Bantul",
				}
				newUserByte, _ := json.Marshal(newUser)
				fmt.Fprintf(writer, string(newUserByte))
			}
			var server = http.Server{
				Addr:    "localhost:8081",
				Handler: handler,
			}
			if err := server.ListenAndServe(); err != nil {
				panic(err)
			}
		}()

		/* Server 3 */
		go func() {
			var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("Test 1..2..3.."))
			}
			var server = http.Server{
				Addr:    "localhost:8082",
				Handler: handler,
			}
			if err := server.ListenAndServe(); err != nil {
				panic(err)
			}
		}()

	}()
	<-forever
}
