// package main

// import "fmt"

// type Person struct {
// 	FirstName string
// 	LastName  string
// }

// func main() {
// 	name := "Bill"
// 	fmt.Println(*&name)

// 	mySLice := []string{"Madrid", "Tokyo", "London"}
// 	fmt.Println(mySLice)
// 	updateSlice(mySLice)
// 	fmt.Println(mySLice)

// 	person := Person{"Noval", "Wardhana"}
// 	fmt.Println(person)
// 	person.UpdateName("Dhana")
// 	fmt.Println(person)
// }

// func (p *Person) UpdateName(name string) {
// 	(*p).FirstName = name
// }

// func updateSlice(slice []string) {
// 	slice[0] = "Bantul"
// }

package main

import (
	"fmt"
)

func main() {

	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(namePointer)
}
