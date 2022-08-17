package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	FirstName string `required:"true" max:"10"`
	LastName  string `required:"true" max:"11" gorm:"-"`
	Age       int    `required:"true" max:"12" gorm:"true"`
}

func isValid(data interface{}) bool {
	dataType := reflect.TypeOf(data)
	for i := 0; i < dataType.NumField(); i++ {
		if dataType.Field(i).Tag.Get("required") != "true" {
			if reflect.ValueOf(data).Field(i).Interface() == "" {
				return false
			}
		}
	}

	return true
}

func main() {
	/*
		Reflect
		1. Dalam bahasa pemrograman biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode kita saat aplikasi sedang berjalan
		2. Reflect sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
	*/
	person := Person{"Noval", "Wardhana", 28}
	personType := reflect.TypeOf(person)
	personTypeField := personType.Field(0)
	fmt.Println(personTypeField.Name)
	personTypeField = personType.Field(1)
	fmt.Println(personTypeField.Name)
	personTypeField = personType.Field(2)
	fmt.Println(personTypeField.Name)

	fmt.Println(personType.Field(0).Tag.Get("required"))
	fmt.Println(personType.Field(0).Tag.Get("max"))
	fmt.Println(personType.Field(1).Tag.Get("required"))
	fmt.Println(personType.Field(1).Tag.Get("max"))
	fmt.Println(personType.Field(1).Tag.Get("gorm"))
	fmt.Println(personType.Field(2).Tag.Get("gorm"))

	/* Membuat library */
	person = Person{
		FirstName: "Noval",
		Age:       28,
	}
	result := isValid(person)
	fmt.Println(result)
	person = Person{
		FirstName: "Novalita",
		LastName:  "Wardhana",
		Age:       29,
	}
	result = isValid(person)
	fmt.Println(result)
}
