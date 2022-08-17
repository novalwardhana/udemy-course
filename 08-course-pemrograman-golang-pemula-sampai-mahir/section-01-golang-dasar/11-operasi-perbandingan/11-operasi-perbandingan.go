package main

import "fmt"

func main() {
	var nameOne = "Noval"
	var nameTwo = "Noval"
	var nameThree = "noval"

	var result bool = nameOne == nameTwo
	fmt.Println(result)
	result = nameTwo == nameThree
	fmt.Println(result)

	value1 := 100
	value2 := 300
	value3 := 100.00
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
	fmt.Println(value1 == int(value3))
	fmt.Println(float64(value1) == value3)
}
