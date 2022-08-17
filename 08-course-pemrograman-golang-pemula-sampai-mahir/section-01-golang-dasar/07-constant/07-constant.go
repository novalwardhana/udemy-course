package main

import "fmt"

func Constant() {
	const firstName string = "Novalita"
	const midName = "Kusuma"
	const lastName = "Wardhana"
	const age = 28
	fmt.Println(firstName)
	fmt.Println(age)

	const (
		addressOne = "Bantul"
		addressTwo = "Kota Jogja"
	)
	fmt.Println(addressOne, addressTwo)
}
