package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Marriage   bool
	Hobbies    []string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Noval",
		MiddleName: "Kusuma",
		LastName:   "Wardhana",
		Age:        29,
		Marriage:   false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

// TestJSONObjectWithGoroutine:
func TestJSONObjectWithGoroutine(t *testing.T) {

	type User struct {
		FullName string
		Email    string
		Address  string
		Age      int64
	}
	var users = []User{
		{"Noval Wardhana", "noval@gmail.com", "Bantul", 29},
		{"Deni Eka Nugraha", "deni@gmail.com", "Sleman", 30},
		{"Noval Arif Kardianto", "noval@gmail.com", "Sleman", 30},
		{"Adrie Satrio", "adrie@gmail.com", "Bengkulu", 31},
	}
	var userParam = make(chan User)

	var customers = []Customer{
		{FirstName: "Cristiano", LastName: "Ronaldo"},
		{FirstName: "Lionel", LastName: "Messi"},
	}
	var customerParam = make(chan Customer)

	for _, user := range users {
		data := user
		go func() {
			userParam <- data
		}()
	}
	for _, customer := range customers {
		data := customer
		go func() {
			customerParam <- data
		}()
	}

	go func() {
		for {
			select {
			case data := <-userParam:
				{
					bytes, err := json.Marshal(data)
					if err != nil {
						break
					}
					fmt.Println("User data:", string(bytes))
				}
			case data := <-customerParam:
				{
					bytes, err := json.Marshal(data)
					if err != nil {
						panic(err)
					}
					fmt.Println("Customer data:", string(bytes))
				}
			}
		}
	}()

	time.Sleep(10 * time.Second)

}
