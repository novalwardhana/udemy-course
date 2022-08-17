package main

import "fmt"

func main() {
	var clubs [4]string
	clubs[0] = "Manchester United"
	clubs[1] = "Arsenal"
	clubs[2] = "Chelsea"
	clubs[3] = "Liverpool"
	fmt.Println(clubs)
	fmt.Println(len(clubs))

	var values = [5]int{
		2,
		4,
		8,
		16,
		32,
	}
	fmt.Println(values)
	fmt.Println(len(values))

	var values2 [12]float64
	values2[2] = 30
	values2[8] = 45.55
	fmt.Println(values2)
	fmt.Println(len(values2))
	fmt.Println(values2[9])
	fmt.Println(values2[11])
}
