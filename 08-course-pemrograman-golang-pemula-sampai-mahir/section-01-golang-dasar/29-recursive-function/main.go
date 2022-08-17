package main

import "fmt"

func main() {
	/*
		Recursice function
		FUncton yang memanggil function dirinya sendiri
	*/
	fact := factorial(5)
	fmt.Println(fact)
	fact = factorial(10)
	fmt.Println(fact)
	fact = factorial(20)
	fmt.Println(fact)
}

func factorial(number int) int {
	if number <= 1 {
		return 1
	}
	return number * factorial(number-1)
}
