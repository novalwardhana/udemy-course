package main

import (
	"errors"
	"fmt"
)

func divide(value1, value2 int) (int, error) {
	if value2 == 0 {
		return 0, errors.New("divided by zero")
	} else {
		return value1 / value2, nil
	}
}

func main() {
	/*
		Error Interface
		Golang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interfacenya adalah error

		Membuat Error
		1. Untuk membuat error, kita tidak perlu manual
		2. Golang sudah menyediakan package errors
	*/
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("result: ", result)
	}
	result, err = divide(10, 5)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("result: ", result)
	}
}
