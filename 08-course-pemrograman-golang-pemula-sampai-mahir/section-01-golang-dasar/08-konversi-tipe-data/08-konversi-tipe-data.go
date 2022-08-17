package main

import "fmt"

func KonversiTipeData() {
	var number8 int8 = 120
	var number16 int16 = int16(number8)
	var number32 int32 = int32(number16)
	var number64 int64 = int64(number32)
	fmt.Println(number8, number16, number32, number64)

	var fullName string = "Noval Wardhana"
	var splitName byte = fullName[0]
	fmt.Println(string(splitName))
}
