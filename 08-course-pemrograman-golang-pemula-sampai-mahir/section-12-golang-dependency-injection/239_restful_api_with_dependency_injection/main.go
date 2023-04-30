package main

import (
	"restful-api/helper"
	run_wire "restful-api/run-wire"
)

func main() {

	forever := make(chan bool)
	go func() {

		go func() {
			server := run_wire.InitializedCategoryServer()
			err := server.ListenAndServe()
			helper.CategoryPanicIfError(err)
		}()

	}()
	<-forever

}
