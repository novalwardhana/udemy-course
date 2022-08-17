package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println(hasName.GetName())
}

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

type Animal struct {
	Name string
}

func (a Animal) GetName() string {
	return a.Name
}

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (c Customer) GetName() string {
	return c.Name
}

func (c Customer) GetAddress() string {
	return c.Address
}

func (c Customer) GetAge() int {
	return c.Age
}

func main() {
	/*
		Interface
		1. Interface adlah tipe data abstract, dia tidak memiliki implementasi langsung
		2. Sebuah interface berisikan definisi-definisi method
		3. Biasanya interface digunakan sebagai kontrak
		Interface adalah tipe data abstract yang didalamnya berisikan definisi-definisi method

		Implementasi Interface
		1. Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
		2. Sehingga kita tidak perlu mengimplementasikan interface secara manual
		3. Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan
		   secara eksplisit akan menggunakan interface yang mana
	*/
	var noval = Person{
		Name: "Novalita Kusuma Wardhana",
	}
	sayHello(noval)

	var animal = Animal{
		Name: "Harimau",
	}
	sayHello(animal)

	var customer = Customer{
		Name:    "Kusuma",
		Address: "Yogyakarta",
		Age:     28,
	}
	sayHello(customer)
}
