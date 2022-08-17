package main

import "fmt"

func Variable() {
	var name string
	name = "Novalita"
	fmt.Println(name)
	name = "Noval Wardhana"
	fmt.Println(name)

	var fullname = "Novalita Kusuma Wardhana"
	fmt.Println(fullname)

	var age = 28
	fmt.Println(age)

	var number = 189.65
	fmt.Println(number)

	address := "Bantul"
	fmt.Println(address)
	address = "Bantul, Yogyakarta"
	fmt.Println(address)

	var (
		firstName = "Novalita"
		midName   = "Kusuma"
		lastName  = "Wardhana"
	)
	fmt.Println(firstName)
	fmt.Println(midName)
	fmt.Println(lastName)
}
