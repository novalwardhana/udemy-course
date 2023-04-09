package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Noval",
		MiddleName: "Kusuma",
		LastName:   "Wardhana",
		Age:        29,
		Hobbies:    []string{"Math", "Programming", "Running"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestDecodeJSONArray(t *testing.T) {
	jsonString := `{"FirstName":"Noval","MiddleName":"Kusuma","LastName":"Wardhana","Age":29,"Marriage":false,"Hobbies":["Math","Programming","Running"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	if err := json.Unmarshal(jsonBytes, customer); err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

type Customer2 struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Marriage   bool
	Hoobies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestComplexJSONArray(t *testing.T) {
	customer := &Customer2{
		FirstName: "Noval",
		Addresses: []Address{
			{
				Street:     "Madrid",
				Country:    "Spain",
				PostalCode: "123",
			}, {
				Street:     "Manchester",
				Country:    "England",
				PostalCode: "456",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestDecodeComplexJSONArray(t *testing.T) {
	jsonString := `{"FirstName":"Noval","MiddleName":"","LastName":"","Age":0,"Marriage":false,"Hoobies":null,"Addresses":[{"Street":"Madrid","Country":"Spain","PostalCode":"123"},{"Street":"Manchester","Country":"England","PostalCode":"456"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer2{}
	if err := json.Unmarshal(jsonBytes, customer); err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestDecodeArrayJSON(t *testing.T) {
	jsonString := `[{"Street":"Madrid","Country":"Spain","PostalCode":"123"},{"Street":"Manchester","Country":"England","PostalCode":"456"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	if err := json.Unmarshal(jsonBytes, addresses); err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}

func TestEncodeArrayJSON(t *testing.T) {
	Addresses := []Address{
		{
			Street:     "Jakarta",
			Country:    "Indonesia",
			PostalCode: "55000",
		}, {
			Street:     "Banguntapan",
			Country:    "Indonesia",
			PostalCode: "55198",
		},
	}

	bytes, _ := json.Marshal(Addresses)
	fmt.Println(string(bytes))
}

// TestMultipleEncodeJSONArray:
func TestMultipleEncodeJSONArray(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(3)

	type Team struct {
		Name    string
		Code    string
		Country string
		Since   int64
	}
	type FootballPlayer struct {
		FirstName string
		LastName  string
		Age       int64
		Teams     []Team
	}

	/* Encode 1 */
	go func() {
		defer wg.Done()

		var datas = []interface{}{1, 2, "3", true, false, nil, 80}
		bytes, err := json.Marshal(datas)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	}()

	/* Encode 2 */
	go func() {
		defer wg.Done()

		var player = []FootballPlayer{
			{
				FirstName: "Karim",
				LastName:  "Benzema",
				Age:       34,
				Teams: []Team{
					{Name: "Lyon FC", Code: "LFC", Country: "France", Since: 1870},
					{Name: "Real Madrid CF", Code: "RMA", Country: "Spain", Since: 1780},
				},
			},
			{
				FirstName: "Cristiano",
				LastName:  "Ronaldo",
				Age:       37,
				Teams: []Team{
					{Name: "Sporting Lisbon", Code: "SL", Country: "Portugal", Since: 2000},
					{Name: "Manchester United", Code: "MU", Country: "England", Since: 2002},
					{Name: "Real Madrid", Code: "RMA", Country: "Spain", Since: 2009},
				},
			},
		}
		bytes, err := json.Marshal(player)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	}()

	/* Encode 3 */
	go func() {
		defer wg.Done()

		var teams = []Team{
			{Name: "Real Madrid CF", Code: "RMA", Country: "Spain", Since: 1902},
			{Name: "Barcelona FC", Code: "Barca", Country: "Spain", Since: 1899},
		}
		bytes, err := json.Marshal(teams)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	}()

	wg.Wait()
	fmt.Println("Finish")
}

// TestMultipleDecodeJSONArray:
func TestMultipleDecodeJSONArray(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(3)

	type Team struct {
		Name    string
		Code    string
		Country string
		SInce   int64
	}
	type FootballPlayer struct {
		FirstName string
		LastName  string
		Age       int64
		Teams     []Team
	}

	/* Decode 1 */
	go func() {
		defer wg.Done()

		var teams []Team
		var arrayString = `[{"Name":"Real Madrid CF","Code":"RMA","Country":"Spain","Since":1902},{"Name":"Barcelona FC","Code":"Barca","Country":"Spain","Since":1899}]`
		if err := json.Unmarshal([]byte(arrayString), &teams); err != nil {
			panic(err)
		}
		fmt.Println(teams)

	}()

	/* Decode 2 */
	go func() {
		defer wg.Done()

		var footballPlayers []FootballPlayer
		var arrayString = `[{"firstname":"Karim","lastname":"Benzema","Age":34,"Teams":[{"Name":"Lyon FC","Code":"LFC","Country":"France","Since":1870},{"Name":"Real Madrid CF","Code":"RMA","Country":"Spain","Since":1780}]},{"firsTNAme":"Cristiano","lasTName":"Ronaldo","Age":37,"Teams":[{"Name":"Sporting Lisbon","Code":"SL","Country":"Portugal","Since":2000},{"Name":"Manchester United","Code":"MU","Country":"England","Since":2002},{"Name":"Real Madrid","Code":"RMA","Country":"Spain","Since":2009}]}]`
		if err := json.Unmarshal([]byte(arrayString), &footballPlayers); err != nil {
			panic(err)
		}
		fmt.Println(footballPlayers)
	}()

	/* Decode 3 */
	go func() {
		defer wg.Done()

		var arrs []interface{}
		var arrayString = `[1,2,"3",true,false,null,80]`
		if err := json.Unmarshal([]byte(arrayString), &arrs); err != nil {
			panic(err)
		}
		fmt.Println(arrs...)

	}()

	wg.Wait()
	fmt.Println("Finish")
}
