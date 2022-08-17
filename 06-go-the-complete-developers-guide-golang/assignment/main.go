package main

import "fmt"

func main() {
	params := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	EvenOddChecker(params)
}

func EvenOddChecker(datas []int) {
	for _, data := range datas {
		if data%2 == 0 {
			fmt.Println(data, " is even")
		} else {
			fmt.Println(data, " is odd")
		}
	}
}
