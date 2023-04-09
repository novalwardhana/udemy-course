package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func logJSON(param interface{}) {
	bytes, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

// TestEncodeJSON:
func TestEncodeJSON(t *testing.T) {
	logJSON("noval")
	logJSON(1)
	logJSON(true)
	logJSON([]string{"Novalita", "Kusuma", "Wardhana"})
}

// TestEncodeJSONWithGoroutine:
func TestEncodeJSONWithGoroutine(t *testing.T) {

	var encode = func(param interface{}) ([]byte, error) {
		bytes, err := json.Marshal(param)
		return bytes, err
	}
	var param = make(chan interface{})
	var params = []interface{}{1, true, "noval", "wardhana", Customer{FirstName: "Noval", MiddleName: "Kusuma", LastName: "Wardhana"}}

	for _, data := range params {
		newParam := data
		go func() {
			param <- newParam
		}()
	}

	go func() {
		for {
			select {
			case data := <-param:
				{
					bytes, err := encode(data)
					if err != nil {
						break
					}
					fmt.Println("Result: ", string(bytes))
				}
			}
		}
	}()
	time.Sleep(10 * time.Second)

	fmt.Println("Finish")

}
