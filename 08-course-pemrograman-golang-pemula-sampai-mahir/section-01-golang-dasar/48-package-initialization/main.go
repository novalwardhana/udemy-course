package main

import (
	_ "48-package-initialization/mysql"
	"48-package-initialization/postgresql"
	"fmt"
)

func main() {
	/*
		Package Initialization
		1. Saat kita membuat package, kita bisa membuat sebuah function yang diakses ketika package kita diakses
		2. Ini sanagt cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database,
		   Kita bisa membuat function initialisasi database
		3. Kita cukup membuat package dengan nama init

		Blank Identifier
		1. Kadang kita hanya ingin menjalankan init function
		2. Secara default golang akan komplain ketika import package tetapi tidak dijalankan
		3. Untuk menangai hal tersebut, kita bisa menggunakan "_" sebelum nama package ketika import
	*/
	fmt.Println(postgresql.GetDatabase())
	fmt.Println(postgresql.GetPort())
	fmt.Println(postgresql.GetHost())
}
