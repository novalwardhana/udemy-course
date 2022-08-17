package main

import "fmt"

func main() {
	/*
		Variadic function: Jika paramater di dalam sebuah function memiliki tipe data yang sama
	*/
	sum := sumValue(3.5, 6.2, 9.8, 7.2, 5.9, 8.7, 4.9, 7.9)
	fmt.Println(sum)
	max, min := maxMinValue(3.5, 6.2, 9.8, 7.2, 5.9, 8.7, 4.9, 7.9)
	fmt.Println(max)
	fmt.Println(min, "\n")

	/* Variadic function with slice */
	params := []float64{1.5, 2.5, 3.5, 4.5, 5.5}
	sum = sumValue(params...)
	max, min = maxMinValue(params...)
	fmt.Println(sum)
	fmt.Println(max)
	fmt.Println(min, "\n")

	sum = sumValue([]float64{2.6, 3.6, 4.6, 5.6, 6.6, 7.6}...)
	max, min = maxMinValue([]float64{2.6, 3.6, 4.6, 5.6, 6.6, 7.6}...)
	fmt.Println(sum)
	fmt.Println(max)
	fmt.Println(min)
}

func sumValue(datas ...float64) float64 {
	var sum float64
	for _, data := range datas {
		sum += data
	}
	return sum
}

func maxMinValue(datas ...float64) (max float64, min float64) {
	max = 0
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
