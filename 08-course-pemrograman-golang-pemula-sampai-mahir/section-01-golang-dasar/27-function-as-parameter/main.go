package main

import "fmt"

type maxData func(...int) int
type minData func(...int) int
type sumData func(...int) int

func main() {
	/*
		Function as Parameter
		1. Function tidak hanya bisa kita simpan di dalam variabel sebagai value
		2. Namun juga bisa kita gunakan sebagai parameter untuk function lain
	*/
	statistic([]int{9, 8, 3, 5, 2, 6, 9, 11, 6, 8}, getMax, getMin, getSum)

	/* Function as parameter dengan type declaration */
	statisticAlias([]int{9, 8, 3, 5, 2, 6, 9, 11, 6, 8}, getMax, getMin, getSum)
}

func statistic(datas []int, max func(...int) int, min func(...int) int, sum func(...int) int) {
	maxData := max(datas...)
	minData := min(datas...)
	sumData := sum(datas...)
	fmt.Println("max: ", maxData)
	fmt.Println("min: ", minData)
	fmt.Println("sum: ", sumData)
}

func statisticAlias(datas []int, max maxData, min minData, sum sumData) {
	maxData := max(datas...)
	minData := min(datas...)
	sumData := sum(datas...)
	fmt.Println(maxData)
	fmt.Println(minData)
	fmt.Println(sumData)
}

func getMax(datas ...int) int {
	var max int
	for _, data := range datas {
		if max < data {
			max = data
		}
	}
	return max
}

func getMin(datas ...int) int {
	var min int = datas[0]
	for _, data := range datas {
		if min > data {
			min = data
		}
	}
	return min
}

func getSum(datas ...int) int {
	var sum int
	for _, data := range datas {
		sum += data
	}
	return sum
}
