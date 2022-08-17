package main

import "fmt"

func main() {
	first, middle, last := getFullName()
	fmt.Println(first)
	fmt.Println(middle)
	fmt.Println(last)

	max, min, length := maxMinLengthNumber([]float64{3.1, 2.3, 8.99, 0.24, 9.25, 7.45, 8.25})
	fmt.Println(max)
	fmt.Println(min)
	fmt.Println(length)
}

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Novalita"
	middleName = "Kusuma"
	lastName = "Wardhana"
	return
}

func maxMinLengthNumber(datas []float64) (max float64, min float64, length int) {
	length = len(datas)
	if length == 0 {
		return
	}
	min = datas[0]
	for _, data := range datas {
		if max < data {
			max = data
		}
		if min > data {
			min = data
		}
	}
	return
}
