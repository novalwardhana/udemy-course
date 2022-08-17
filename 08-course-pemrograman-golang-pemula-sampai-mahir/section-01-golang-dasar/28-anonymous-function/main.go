package main

import "fmt"

type maxData func([]float64) float64
type minData func([]float64) float64
type sumData func([]float64) float64
type meanDta func([]float64) float64

func main() {
	/*
		Anonymous Function
		Membuat function langsung di variabel atau parameter tanpa harus membuat function terlebih dahulu
	*/
	max := func(datas []float64) float64 {
		var max float64
		for _, data := range datas {
			if max < data {
				max = data
			}
		}
		return max
	}
	min := func(datas []float64) float64 {
		var min float64 = datas[0]
		for _, data := range datas {
			if min > data {
				min = data
			}
		}
		return min
	}
	statistic([]float64{6.7, 3.2, 9.8, 1.6, 4.7, 8.2, 11.5, 5.5, 8.6, 9.8}, max, min,
		func(datas []float64) float64 {
			var sum float64
			for _, data := range datas {
				sum += data
			}
			return sum
		},
		func(datas []float64) float64 {
			var sum float64
			for _, data := range datas {
				sum += data
			}
			return sum / float64(len(datas))
		},
	)

	sum := func(datas []float64) float64 {
		var sum float64
		for _, data := range datas {
			sum += data
		}
		return sum
	}
	mean := func(datas []float64) float64 {
		sum := sum(datas)
		return sum / float64(len(datas))
	}
	statistic([]float64{10.11, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9, 1.2, 9.1},
		func(datas []float64) float64 {
			var max float64
			for _, data := range datas {
				if max < data {
					max = data
				}
			}
			return max
		},
		func(datas []float64) float64 {
			var min float64
			for _, data := range datas {
				if min > data {
					min = data
				}
			}
			return min
		},
		sum,
		mean,
	)
}

func statistic(datas []float64, max maxData, min minData, sum sumData, mean meanDta) {
	getMax := max(datas)
	getMin := min(datas)
	getSum := sum(datas)
	getMean := mean(datas)
	fmt.Println("getMax: ", getMax)
	fmt.Println("getMin: ", getMin)
	fmt.Println("getSum: ", getSum)
	fmt.Println("getMean: ", getMean)
}
