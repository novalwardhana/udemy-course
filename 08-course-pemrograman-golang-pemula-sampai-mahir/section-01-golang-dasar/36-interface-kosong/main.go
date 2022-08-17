package main

import (
	"errors"
	"fmt"
)

type apapun interface{}
type bebas interface{}

func TestInterface(param int, apapun interface{}, bebas interface{}) interface{} {
	if param == 1 {
		return 1
	} else if param == 2 {
		return true
	} else if param == 3 {
		return 10.5
	} else if param == 4 {
		return errors.New("Test try error")
	} else if param == 5 {
		panic("Test panic")
		return "Test panic"
	} else {
		return "test"
	}
}

func main() {
	/*
		Interface Kosong
		1. Golang bukanlah bahasa pemrograman berorientasi objek
		2. Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada
		   di bahasa pemrograman tersebut
		3. Contoh di Java ada java.lang.Object
		4. Untuk menangani kasus tersebut Golang menggunakan interface kosong
		5. Interface kosong adalah interface yang tidak memiliki deklarasi method satupun,
		   hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya

		Penggunaan inteface kosong
		1. fmt.Println(a ...interface{})
		2. panic(v interface{})
		3. recover() interface{}
		4. dan lain-lain
	*/

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Test recover", r)
		}
	}()

	var intfOne apapun
	var intfTwo bebas
	var result = TestInterface(1, intfOne, intfTwo)
	fmt.Println(result)
	result = TestInterface(2, intfOne, intfTwo)
	fmt.Println(result)
	result = TestInterface(3, intfOne, intfTwo)
	fmt.Println(result)
	result = TestInterface(4, intfOne, intfTwo)
	fmt.Println(result)
	//result = TestInterface(5, intfOne, intfTwo)
	//fmt.Println(result)
	result = TestInterface(6, intfOne, intfTwo)
	fmt.Println(result)
}
