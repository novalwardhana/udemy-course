package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Function value
		1. Function adalah first class citizen
		2. Function merupakan tipe data dan bisa disimpan dalam bentuk variabel
	*/

	getSayHello := sayHello
	trySayHello := sayHello
	fmt.Println(getSayHello) // Memiliki momory yang sama
	fmt.Println(trySayHello) // Memiliki memory yang sama

	result := getSayHello("Noval")
	fmt.Println(result)
	result = trySayHello("Wardhana")
	fmt.Println(result)

	newHypotenuse := getHypotenuse
	fmt.Println(newHypotenuse(3, 4))
	fmt.Println(newHypotenuse(6, 8))
}

func sayHello(name string) string {
	return "Hallo" + name
}

func getHypotenuse(length, wide float64) (result float64) {
	result = math.Sqrt((length*length + wide*wide))
	return
}
