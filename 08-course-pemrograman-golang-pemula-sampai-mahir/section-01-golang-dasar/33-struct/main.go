package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Marriage      bool
}

func main() {
	/*
		Struct
		1. Struct adalah template data untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
		2. Struct biasanya adalah representasi data dalam program aplikasi yang kita buat
		3. Data struct disimpan dalam field
		4. Sederhananya struct adalah kumpulan dari field

		Membuat Data Struct
		1. Struct adalah template data atau prototype data
		2. Struct tidak bisa langsung digunakan
		3. Namun kita bisa membuat data/object dari struct yang telah kita buat
	*/

	var noval Customer
	noval.Name = "Novalita Kusuma Wardhana"
	noval.Age = 28
	noval.Address = "Bantul"
	fmt.Println(noval.Name)
	fmt.Println(noval.Address)
	fmt.Println(noval.Age)

	/*
		Struct Literals
	*/
	var kusuma = Customer{
		Name:     "Kusuma",
		Address:  "Banguntapan",
		Age:      28,
		Marriage: false,
	}
	fmt.Println(kusuma)
	wardhana := Customer{"Wardhana", "Maguwo", 29, true}
	fmt.Println(wardhana)
}
