package main

import "fmt"

func main() {
	sayMyName("Noval", "wardhana")
	firstName := "Dhana"
	lastName := "Kusuma"
	sayMyName(firstName, lastName)
}

func sayMyName(firstName string, lastName string) {
	fmt.Println("Hallo", firstName, lastName)
}
