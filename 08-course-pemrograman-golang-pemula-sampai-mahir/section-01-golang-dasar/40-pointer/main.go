package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/*
		Pass by Value
		1. Secara default di golang semua variable di passing by value, bukan reference
		2. Artinya, jika kita mengirim sebuah variabel ke dalam function, method atau variabel lain,
		   sebenarnya yang dikirim adalah duplikasi value nya
	*/
	adress1 := Address{"Bantul", "Yogyakarta", "Indonesia"}
	adress2 := adress1
	adress2.City = "Sleman"
	fmt.Println(adress1)
	fmt.Println(adress2)

	/*
		Pointer
		1. Pointer adalah kemapuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
		2. Sederhanya dengan kemapuan pointer, kita bisa membuat pass by reference

		Operator &
		1. Untuk membuat sebuah variabel dengan nilai pointer ke variabel lain, kita bisa menggunakan operator & diikuti nama variabel nya
	*/
	var address3 Address = Address{"Gunungkidul", "Yogyakarta", "Indonesia"}
	var address4 *Address = &address3
	address4.City = "Kulon Progo"
	fmt.Println(address3)
	fmt.Println(address4)

	/*
		Operator *
		1.Saat kita mengubah variabel pointer, maka yang berubah hanya variabel tersebut
		2.Semua variabel yang mengacu data yang sama tidak akan berubah
		3.Jika kita ingin mengubah seluruh variabel yang mengacu pada data tersebut, kita bisa menggunakan operator *
	*/
	/* Contoh variabel tidak berubah */
	var address5 Address = Address{"Solo", "Jawa Tengah", "Indonesia"}
	var address6 *Address = &address5
	address6.City = "Klaten"
	address6 = &Address{"Magetan", "Jawa Timur", "Indonesia"}
	fmt.Println(address5)
	fmt.Println(address6)

	/* Contoh variabel berubah */
	var address7 Address = Address{"Salatiga", "Jawa Tengah", "Indonesia"}
	var address8 *Address = &address7
	var address9 *Address = address8
	address8.City = "Magelang"
	*address8 = Address{"Bandung", "Jawa Barat", "Indonesia"}
	fmt.Println(address7)
	fmt.Println(address8)
	fmt.Println(address9)

	/*
		Function New
		1. Golang juga memiliki function new yang bisa digunakan untuk membuat pointer
		2. Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
	*/
	var address10 *Address = new(Address)
	var address11 *Address = address10
	address11.City = "Jakarta Selatan"
	address10.Province = "Daerah Khusus Ibukota Jakarta"
	address11.Country = "Indonesia"
	fmt.Println(address10)
	fmt.Println(address11)
}
