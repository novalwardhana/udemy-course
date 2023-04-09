package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageURL string `json:"image_url"`
}

// TestJSONTag:
func TestJSONTag(t *testing.T) {
	product := Product{
		ID:       "P001",
		Name:     "Iphone 14",
		Price:    12500000,
		ImageURL: "apple.com",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

// TestDecodeJSONTag:
func TestDecodeJSONTag(t *testing.T) {
	jsonString := `{"Id":"P001","Name":"Iphone 14","pRICe":12500000,"ImageUrl":"apple.com"}` // tidak case sensitive
	jsonBytes := []byte(jsonString)

	Product := &Product{}
	json.Unmarshal(jsonBytes, Product)
	fmt.Println(Product)
}

type User struct {
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Address   UserAddress `json:"address"`
	Roles     []Role      `json:"roles"`
}

type UserAddress struct {
	Name       string `json:"name"`
	PostalCode int64  `json:"postal_code"`
}

type Role struct {
	Name string `json:"name"`
}

// TestMultipleJSONTag:
func TestEncodeJSONTagWithGoroutine(t *testing.T) {
	users := []User{
		{
			FirstName: "Noval",
			LastName:  "Wardhana",
			Email:     "noval@gmail.com",
			Address: UserAddress{
				Name:       "Banguntapan",
				PostalCode: 55198,
			},
			Roles: []Role{
				{Name: "Admin"},
				{Name: "Root"},
			},
		},
		{
			FirstName: "Eko Kurniawan",
			LastName:  "Khannedy",
			Email:     "echo@programmerzamannow.com",
			Address: UserAddress{
				Name:       "Bandung",
				PostalCode: 55199,
			},
			Roles: []Role{
				{Name: "Admin"},
				{Name: "Guest"},
			},
		},
	}
	var param = make(chan User)

	for _, user := range users {
		data := user
		go func() {
			param <- data
		}()
	}

	go func() {
		for {
			select {
			case data := <-param:
				{
					bytes, err := json.Marshal(data)
					if err != nil {
						break
					}
					fmt.Println(string(bytes))
				}
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Finish")
}

// TestDecodeJSONTagWithGoroutine:
func TestDecodeJSONTagWithGoroutine(t *testing.T) {
	var userStrings = []string{
		`{"FIRST_name":"Noval","LAST_name":"Wardhana","email":"noval@gmail.com","address":{"name":"Banguntapan","POSTAL_code":55198},"roles":[{"name":"Admin"},{"name":"Root"}]}`,
		`{"first_NAME":"Eko Kurniawan","last_NAME":"Khannedy","email":"echo@programmerzamannow.com","address":{"name":"Bandung","postal_CODE":55199},"roles":[{"name":"Admin"},{"name":"Guest"}]}`,
	}
	var param = make(chan string, 2)

	for _, userString := range userStrings {
		data := userString
		go func() {
			param <- data
		}()
	}

	go func() {
		for {
			select {
			case data := <-param:
				{
					var user = new(User)
					if err := json.Unmarshal([]byte(data), user); err != nil {
						break
					}
					fmt.Println(user)
				}
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Finish")

}
