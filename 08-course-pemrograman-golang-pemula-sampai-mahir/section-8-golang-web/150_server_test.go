package belajar_golang_web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:9000",
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestServerAdvance(t *testing.T) {
	forever := make(chan int)
	for _, address := range []string{"localhost:9001", "localhost:9002", "localhost:9003"} {
		addr := address
		go func() {
			server := http.Server{
				Addr: addr,
			}
			if err := server.ListenAndServe(); err != nil {
				panic(err)
			}
		}()
	}
	<-forever
}
