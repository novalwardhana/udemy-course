package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*
		Container List
		Adalah implementasi struktur data double linked list di golang
	*/

	data := list.New()
	data.PushBack("Novalita")
	data.PushBack("Kusuma")
	data.PushBack("Wardhana")
	data.PushFront("Mr.")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

	data = list.New()
	data.PushBack("Arsenal")
	data.PushBack("Liverpool")
	data.PushBack("Tottenham")
	data.PushFront("Manchester United")
	data.PushFront("Chelsea")
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
