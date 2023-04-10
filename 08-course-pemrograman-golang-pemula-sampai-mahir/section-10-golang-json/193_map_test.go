package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

// TestMapDecode:
func TestMapDecode(t *testing.T) {
	jsonRequest := `{"id": "123456", "name": "Macbook Pro M1 Pro", "price": 25000000}`
	jsonByte := []byte(jsonRequest)

	var result map[string]interface{}
	if err := json.Unmarshal(jsonByte, &result); err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

// TestMapEncode:
func TestMapEncode(t *testing.T) {
	var product = map[string]interface{}{
		"id":    "123456",
		"name":  "Macbook Pro M1 Pro",
		"price": 25000000,
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

// TestMapDecodeWithGoroutine:
func TestMapDecodeWithGoroutine(t *testing.T) {
	var jsonRequests = []string{
		`{"random":"16","random float":"19.662","bool":"true","date":"1999-10-31","regEx":"hello to you","enum":"online","firstname":"Amara","lastname":"Juliet","city":"Taiyuan","country":"Qatar","countryCode":"AM","email uses current data":"Amara.Juliet@gmail.com","email from expression":"Amara.Juliet@yopmail.com"}`,
		`{"random":"81","random float":"15.197","bool":"true","date":"1982-09-21","regEx":"hello to you","enum":"json","firstname":"Monika","lastname":"Even","city":"Seattle","country":"New Zealand","countryCode":"BF","email uses current data":"Monika.Even@gmail.com","email from expression":"Monika.Even@yopmail.com"}`,
		`{"random":"68","random float":"29.721","bool":"true","date":"1984-12-26","regEx":"hello to you","enum":"json","firstname":"Dorthy","lastname":"Rurik","city":"Semarang","country":"Guyana","countryCode":"ZA","email uses current data":"Dorthy.Rurik@gmail.com","email from expression":"Dorthy.Rurik@yopmail.com"}`,
	}
	var param = make(chan string)
	for _, jsonRequest := range jsonRequests {
		data := jsonRequest
		go func() {
			param <- data
		}()
	}

	go func() {
		for {
			select {
			case data := <-param:
				{
					var mapData map[string]interface{}
					if err := json.Unmarshal([]byte(data), &mapData); err != nil {
						break
					}
					fmt.Println(mapData)
				}
			}
		}
	}()

	time.Sleep(5 * time.Second)
}

// TestMapEncodeWithGoroutine:
func TestMapEncodeWithGoroutine(t *testing.T) {

	var param = make(chan interface{})

	/* Data 1 */
	go func() {
		var data = map[string]interface{}{
			"FirstName":  "Noval",
			"MiddleName": "Kusuma",
			"LastNAme":   "Wardhana",
		}
		param <- data
	}()

	/* Data 2 */
	go func() {
		var data = map[string]interface{}{
			"user1": User{
				FirstName: "Noval",
				LastName:  "Wardhana",
				Email:     "noval@gmail.com",
				Address: UserAddress{
					Name:       "Bantul",
					PostalCode: 55198,
				},
				Roles: []Role{
					{Name: "Root"},
					{Name: "Admin"},
				},
			},
			"User2": User{
				FirstName: "Adrie",
				LastName:  "Satrio",
				Email:     "adrie@gmail.com",
				Address: UserAddress{
					Name:       "Bengkulu",
					PostalCode: 212312,
				},
				Roles: []Role{
					{Name: "Guest"},
				},
			},
		}
		param <- data
	}()

	/* Data 3 */
	go func() {
		var data = map[string]interface{}{
			"Customer1": Customer{
				FirstName:  "Deni",
				MiddleName: "Eka",
				LastName:   "Nugraha",
				Age:        30,
				Marriage:   true,
				Hobbies:    []string{"Photography", "Motovlog"},
			},
			"Customer2": Customer2{
				FirstName:  "Novan",
				MiddleName: "Arif",
				LastName:   "Kardianto",
				Age:        30,
				Marriage:   false,
				Hoobies:    []string{"Design", "Photography"},
				Addresses: []Address{
					{Street: "Berbah", Country: "Indonesia", PostalCode: "55199"},
					{Street: "Gunung Kidul", Country: "Indonesia", PostalCode: "212515"},
				},
			},
		}
		param <- data
	}()

	go func() {
		for {
			select {
			case data := <-param:
				{
					bytes, err := json.Marshal(data)
					if err != nil {
						break
					}
					fmt.Println("Map to json: ", string(bytes))
				}
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
