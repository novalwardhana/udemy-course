package main

import "fmt"

func main() {
	person := map[string]string{
		"fistName": "Noval",
		"lastName": "Wardhana",
		"address":  "Bantul",
		"age":      "28",
	}
	person["test"] = "Test data"
	fmt.Println(person)
	fmt.Println(person["address"])
	fmt.Println(len(person))
	delete(person, "test2")
	delete(person, "test")
	fmt.Println(person)
	fmt.Println(len(person))

	personTwo := make(map[string]int)
	personTwo["noval"] = 28
	personTwo["casca"] = 29
	fmt.Println(personTwo)
}
