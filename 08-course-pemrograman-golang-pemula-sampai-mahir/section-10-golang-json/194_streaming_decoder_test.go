package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
)

// TestStreamingDecoder:
func TestStreamingDecoder(t *testing.T) {
	file, err := os.Open("./resources/json_decoder.json")
	if err != nil {
		panic(err)
	}

	var customer Customer
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&customer); err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

// TestCompareStreamingDecoder:
func TestCompareStreamingDecoder(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	/* With package json decoder */
	go func() {
		defer wg.Done()

		file, err := os.Open("./resources/json_decoder_2.json")
		if err != nil {
			panic(err)
		}
		decoder := json.NewDecoder(file)

		type NewCustomer struct {
			Customer1 Customer  `json:"Customer1"`
			Customer2 Customer2 `json:"Customer2"`
		}
		var newCustomer = &NewCustomer{}
		if err := decoder.Decode(newCustomer); err != nil {
			panic(err)
		}
		fmt.Println("With json decoder: ", newCustomer)
	}()

	/* Without package json decoder */
	go func() {
		defer wg.Done()

		file, err := os.Open("./resources/json_decoder_2.json")
		if err != nil {
			panic(err)
		}
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		type NewCustomer struct {
			Customer1 Customer  `json:"Customer1"`
			Customer2 Customer2 `json:"Customer2"`
		}
		var newCustomer = &NewCustomer{}
		if err := json.Unmarshal(bytes, newCustomer); err != nil {
			panic(err)
		}
		fmt.Println("Without json decoder: ", newCustomer)

	}()

	wg.Wait()
	fmt.Println("Finish")
}
