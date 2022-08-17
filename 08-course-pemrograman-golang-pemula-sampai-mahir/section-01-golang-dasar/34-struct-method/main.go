package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

type Cube struct {
	Length float64
	Wide   float64
	High   float64
}

func (c Cube) Volume() float64 {
	return c.Length * c.Wide * c.High
}

func (c Cube) SurfaceArea() (surface float64) {
	surface = 2*c.Length*c.Wide + 2*c.Wide*c.High + 2*c.Length*c.High
	return
}

func main() {
	/*
		Struct Method
		1. Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
		2. Namun jika kita ingin menambahkan method ke dalam structs, sehingga sekakan-akan sebuat struct memiliki function
		3. Function di dalam struct disebut sebagai method
	*/
	var noval = Customer{
		Name:    "Novalita",
		Address: "Bantul",
		Age:     28,
	}
	noval.sayHello("Adrie")
	noval.sayHello("Novan")
	noval.sayHello("Deni")

	cubeOne := Cube{
		Length: 20.5,
		Wide:   5.5,
		High:   25.5,
	}
	fmt.Println("CubeOne:", cubeOne)
	fmt.Println("cubOne volume:", cubeOne.Volume())
	fmt.Println("cubeOne surface area:", cubeOne.SurfaceArea())

	cubeTwo := Cube{30, 10, 50}
	fmt.Println("cubeTwo:", cubeTwo)
	fmt.Println("cubeTwo volume:", cubeTwo.Volume())
	fmt.Println("cubeTwo surface area:", cubeTwo.SurfaceArea())
}
