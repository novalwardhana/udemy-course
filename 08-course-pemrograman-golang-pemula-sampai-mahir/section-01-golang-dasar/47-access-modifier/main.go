package main

import (
	"47-access-modifier/helper"
	"fmt"
)

func main() {
	/*
		Access Modifier
		1. Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menetukan access modifier terhadap suatu function atau variable
		2. DI golang, untuk menentukan access modofier cukup dengan nama function atau variable
		3. Jika namanya diawali dengan huruf besar maka artinya bisa diakses di package lain. Jika dimulai huruf kecil artinya tidak bisa diakses di package lain
	*/

	fmt.Println(helper.Application)
	fmt.Println(helper.Version)
	fmt.Println(helper.SayHello("Wardhana", "Noval"))
}
