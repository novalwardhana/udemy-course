package main

import "fmt"

type Address struct {
	City, Provice, Country string
}

func changeAddress(address Address) {
	address.City = "Bantul"
	address.Provice = "Yogyakarta"
	address.Country = "Indonesia"
}

func changeAddressPointer(address *Address) {
	address.City = "Kulon Progo"
	address.Provice = "Yogyakarta"
	address.Country = "Indonesia"
}

func main() {
	/*
		Pointer Function
		1. Saat kita membuat parameter function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
		2. Oleh karena itu, jika kita ingin mengubah data di dalam function, data yang asli tidak akan berubah
		3. Hal ini membuat variabel aman, karena tidak akan bisa diubah
		4. Namun kadang kita ingin membuat function yang bisa mengubah data asli paramater tersebut
		5. Untuk melakukan ini kita juga bisa menggunakan pointer di function
		6. Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya
	*/
	address1 := Address{"Magetan", "Jawa Timur", "Indonesia"}
	fmt.Println(address1)
	changeAddress(address1)
	fmt.Println(address1)

	address2 := Address{"Ngawi", "Jawa Timur", "Indonesia"}
	fmt.Println(address2)
	changeAddressPointer(&address2)
	fmt.Println(address2)

	address3 := new(Address)
	address3.City = "Cirebon"
	address3.Provice = "Jawa Barat"
	address3.Country = "Indonesia"
	fmt.Println(address3)
	changeAddressPointer(address3)
	fmt.Println(address3)

	address4 := &Address{
		City:    "Surakarta",
		Provice: "Jawa Tengah",
		Country: "Indonesia",
	}
	fmt.Println(address4)
	changeAddressPointer(address4)
	fmt.Println(address4)
}
