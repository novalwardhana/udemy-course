package belajar_golang_web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{}
	server.Addr = "localhost:8080"
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestServerMultiple(t *testing.T) {
	forever := make(chan int)
	go func() {

		/* Server 1 */
		go func() {
			server1 := http.Server{
				Addr: "localhost:8080",
			}
			if err := server1.ListenAndServe(); err != nil {
				panic(err)
			}
		}()

		/* Server 2 */
		go func() {
			server2 := http.Server{}
			server2.Addr = "localhost:8081"
			if err := server2.ListenAndServe(); err != nil {
				panic(err)
			}
		}()

		/* Server 3 */
		go func() {
			server3 := http.Server{
				Addr: "localhost:8082",
			}
			if err := server3.ListenAndServe(); err != nil {
				panic(err)
			}
		}()

	}()
	<-forever
}
