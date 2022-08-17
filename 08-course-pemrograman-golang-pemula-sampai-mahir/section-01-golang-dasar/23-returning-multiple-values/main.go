package main

import "fmt"

func main() {

	/* multiple return balue */
	firstName, lastName := getName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	max, min := maxMinNumber([]int{9, 8, 10, 2, 6, 5, 7, 9, 8, 9})
	fmt.Println(max)
	fmt.Println(min)

	/* Ignore multiple value */
	firstName, _ = getName()
	fmt.Println(firstName)
	max, _ = maxMinNumber([]int{5, 9, 3, 7, 3})
	fmt.Println(max)
}

func getName() (string, string) {
	return "Noval", "Wardhana"
}

func maxMinNumber(datas []int) (int, int) {
	if len(datas) == 0 {
		return 0, 0
	}

	var max int
	var min int = datas[0]
	for _, data := range datas {
		if max < data {
			max = data
		}
		if min > data {
			min = data
		}
	}
	return max, min
}
