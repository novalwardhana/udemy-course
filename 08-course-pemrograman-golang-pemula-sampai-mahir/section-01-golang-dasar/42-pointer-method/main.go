package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (person Person) Married() {
	person.FirstName = "Mr. " + person.FirstName
}

func (person *Person) MarriedPointer() {
	person.FirstName = "Mr. " + person.FirstName
}

type Vertex struct {
	X float64
	Y float64
}

func (v *Vertex) ScaleUp(scale float64) {
	v.X *= scale
	v.Y *= scale
}

func (v *Vertex) ScaleDown(scale float64) {
	v.X /= scale
	v.Y /= scale
}

func main() {
	/*
		Pointer Method
		1.Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses dalam method adalah pass by value
		2.Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena selalu diduplikasi ketika memanggil method
	*/
	person1 := Person{"Noval", "Wardhana", 28}
	fmt.Println(person1)
	person1.Married()
	fmt.Println(person1)
	person1.MarriedPointer()
	fmt.Println(person1)

	person2 := Person{
		FirstName: "Novalita",
		LastName:  "Kusuma Wardhana",
		Age:       28,
	}
	fmt.Println(person2)
	person2.MarriedPointer()
	fmt.Println(person2)

	person3 := &Person{
		FirstName: "Dhana",
		LastName:  "Kusuma",
		Age:       29,
	}
	fmt.Println(person3)
	person3.MarriedPointer()
	fmt.Println(person3)

	vertex1 := Vertex{
		X: 5,
		Y: 15,
	}
	fmt.Println(vertex1)
	vertex1.ScaleDown(10)
	fmt.Println(vertex1)
	vertex1.ScaleUp(100)
	fmt.Println(vertex1)

	vertex2 := Vertex{300, 100}
	fmt.Println(vertex2)
	vertex2.ScaleDown(1000)
	fmt.Println(vertex2)
	vertex2.ScaleUp(3000)
	fmt.Println(vertex2)
}
