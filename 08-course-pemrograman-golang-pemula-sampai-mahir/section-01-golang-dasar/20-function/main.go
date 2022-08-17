package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		sayHai()
	}
	createPattern(10)
}

func sayHai() {
	fmt.Println("Haii world")
}

func createPattern(number int) {
	for i := 1; i <= number; i++ {
		var char string
		for j := 1; j <= number; j++ {
			if i >= j {
				char += "*"
			}
		}
		fmt.Println(char)
	}
}
