package belajar_golang_json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
)

func TestStreamingEncoder(t *testing.T) {
	file, err := os.Create("./resources/json_encoder.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(file)

	customer := Customer{
		FirstName:  "Noval",
		MiddleName: "Kusuma",
		LastName:   "Wardhana",
	}
	if err := encoder.Encode(customer); err != nil {
		panic(err)
	}
}

// TestCompareStreamingEncoder:
func TestCompareStreamingEncoder(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	type NewCustomer struct {
		Customer1 Customer  `json:"customer1"`
		Customer2 Customer2 `json:"customer2"`
	}
	var newCustomer = &NewCustomer{}
	newCustomer.Customer1 = Customer{
		FirstName:  "Deni",
		MiddleName: "Eka",
		LastName:   "Nugraha",
		Age:        30,
		Marriage:   true,
		Hobbies:    []string{"Photography", "Graphic Design"},
	}
	newCustomer.Customer2 = Customer2{
		FirstName:  "Novan",
		MiddleName: "Arif",
		LastName:   "Kardianto",
		Age:        30,
		Marriage:   false,
		Hoobies:    []string{"Web", "Coding"},
		Addresses: []Address{
			{Street: "Jalan berbah", Country: "Indonesia", PostalCode: "55199"},
			{Street: "Jalan Wonosari", Country: "Indonesia", PostalCode: "55200"},
		},
	}

	/* With package json encoder */
	go func() {
		defer wg.Done()

		file, err := os.Create("./resources/json_encoder_2.json")
		if err != nil {
			panic(err)
		}
		encoder := json.NewEncoder(file)

		if err := encoder.Encode(newCustomer); err != nil {
			panic(err)
		}

	}()

	/* Without package json encoder */
	go func() {
		defer wg.Done()

		dataBytes, err := json.Marshal(newCustomer)
		if err != nil {
			panic(err)
		}
		fileSource := io.NopCloser(bytes.NewBuffer(dataBytes))

		file, err := os.Create("./resources/json_encoder_3.json")
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(file, fileSource); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	fmt.Println("Finish")
}
