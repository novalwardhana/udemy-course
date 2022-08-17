package main

import (
	"errors"
	"fmt"
)

func testInterface() interface{} {
	return "Test data"
}

func generateInterface(param int) interface{} {
	if param == 1 {
		return 12
	} else if param == 2 {
		return 12.5
	} else if param == 3 {
		return "test"
	} else if param == 4 {
		return true
	} else if param == 5 {
		return errors.New("Test error")
	} else {
		return "OK"
	}
}

func main() {
	/*
		Type Assertions
		1. Type assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
		2. Fitur ini sering sekali digunakan ketika kita bertemu dengan interface kosong
	*/
	test := testInterface()
	resultTest := test.(string)
	fmt.Println(resultTest)

	/*
		Type Assertions menggunakan Switch
		1. Jika salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
		2. Jika panic tidak terecover, maka otomatis program akan mati
		3. Agar lebih aman sebaiknya kita menggunakan switch expression untuk melakukan type assertions
	*/
	test = generateInterface(3)
	switch test.(type) {
	case int:
		{
			fmt.Println(test, "is integer")
		}
	case float64:
		{
			fmt.Println(test, "is float64")
		}
	case string:
		{
			fmt.Println(test, "is string")
		}
	case error:
		{
			fmt.Println(test, "is error")
		}
	case bool:
		{
			fmt.Println(test, "is bool")
		}
	default:
		{
			fmt.Println(test, "is default string")
		}
	}
}
