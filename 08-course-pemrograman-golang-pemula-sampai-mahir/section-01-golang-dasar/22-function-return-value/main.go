package main

import "fmt"

func main() {
	var result string
	result = sayHello("")
	fmt.Println(result)
	result = sayHello("Noval Wardhana")

	var resultNumber int
	resultNumber = sumNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(resultNumber)
}

func sayHello(name string) string {
	if len(name) > 0 {
		return "Hallo" + name
	} else {
		return "Hallo bro"
	}
}

func sumNumber(datas []int) int {
	var sum int
	for _, data := range datas {
		sum += data
	}
	return sum
}
