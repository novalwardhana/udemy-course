package main

import "fmt"

func printState() {
	fmt.Println("California")

	cards := []string{"a", "b"}
	for _, card := range cards {
		fmt.Println(card)
	}
}
