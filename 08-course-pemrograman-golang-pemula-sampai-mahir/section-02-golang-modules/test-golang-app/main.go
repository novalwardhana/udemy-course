package main

import (
	"fmt"

	test "github.com/novalwardhana/test-golang-module/v2"
	modCube "github.com/novalwardhana/test-golang-module2/cube"
	modStatistic "github.com/novalwardhana/test-golang-module2/statistic"
	modStatistic2 "github.com/novalwardhana/test-golang-module2/v2/statistic"
)

func main() {
	result := test.SayHello("Noval")
	fmt.Println(result)

	newStatistic := modStatistic.Statistic{
		Title: "Food Price List",
		Datas: []float64{12500, 10000.25, 13900.90, 8500.25, 8500.33, 5500, 9000, 7500, 9000, 14500, 15900, 8700, 9500, 6500, 12000},
	}
	var interfaceStatistic modStatistic.StatisticInterface
	interfaceStatistic = &newStatistic
	fmt.Println("Sum: ", interfaceStatistic.Sum())
	fmt.Println("Max: ", interfaceStatistic.Max())
	fmt.Println("Min: ", interfaceStatistic.Min())
	fmt.Println("Mean: ", interfaceStatistic.Mean())

	newCube := modCube.Cube{
		Title:  "My Cube",
		Length: 10.5,
		Wide:   6.8,
		Height: 12.5,
	}
	var interfaceCube modCube.CubeInterface
	interfaceCube = newCube
	fmt.Println("Surface area: ", interfaceCube.SurfaceArea())
	fmt.Println("Volume: ", interfaceCube.Volume())

	newStatistic2 := modStatistic2.Statistic{
		Title: "Computer Price List in USD",
		Datas: []float64{450, 600, 250, 310.45, 700, 400, 230.22, 950, 350, 720},
	}
	var interfaceStatistic2 modStatistic2.StatisticInterface
	interfaceStatistic2 = &newStatistic2
	interfaceStatistic2.Sort()
	fmt.Println("Data: ", interfaceStatistic2.GetData())
	fmt.Println("Sum: ", interfaceStatistic2.Sum())
	fmt.Println("Max: ", interfaceStatistic2.Max())
	fmt.Println("Min: ", interfaceStatistic2.Min())
	fmt.Println("Mean: ", interfaceStatistic2.Mean())
	fmt.Println("Median: ", interfaceStatistic2.Median())
}
