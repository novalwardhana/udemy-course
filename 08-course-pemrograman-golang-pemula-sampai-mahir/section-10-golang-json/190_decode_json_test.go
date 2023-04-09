package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"firstName": "Noval", "middleName": "Kusuma", "lastName": "Wardhana"}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	if err := json.Unmarshal(jsonBytes, customer); err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}

// TestDecodeJSONWithGoroutine:
func TestDecodeJSONWithGoroutine(t *testing.T) {
	var customers = []string{
		`{"FirstName":"Cristiano","MiddleName":"","LastName":"Ronaldo","Age":0,"Marriage":false,"Hobbies":null}`,
		`{"FirstName":"Lionel","MiddleName":"","LastName":"Messi","Age":0,"Marriage":false,"Hobbies":null}`,
	}
	var customerParam = make(chan string, 2)
	for _, customer := range customers {
		data := customer
		go func() {
			customerParam <- data
		}()
	}

	type User struct {
		FullName string
		Email    string
		Address  string
		Age      int64
	}
	var users = []string{
		`{"FullName":"Noval Wardhana","Email":"noval@gmail.com","Address":"Bantul","Age":29}`,
		`{"FullName":"Adrie Satrio","Email":"adrie@gmail.com","Address":"Bengkulu","Age":31}`,
		`{"FullName":"Deni Eka Nugraha","Email":"deni@gmail.com","Address":"Sleman","Age":30}`,
		`{"FullName":"Noval Arif Kardianto","Email":"noval@gmail.com","Address":"Sleman","Age":30}`,
	}
	var userParam = make(chan string, 2)
	for _, user := range users {
		data := user
		go func() {
			userParam <- data
		}()
	}

	go func() {
		for {
			select {
			case data := <-userParam:
				{
					var user = new(User)
					if err := json.Unmarshal([]byte(data), user); err != nil {
						break
					}
					fmt.Println("User data:", user)
				}
			case data := <-customerParam:
				{
					var customer = new(Customer)
					if err := json.Unmarshal([]byte(data), customer); err != nil {
						break
					}
					fmt.Println("Customer data:", customer)
				}
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Finish")

}
