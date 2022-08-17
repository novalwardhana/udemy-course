package main

import (
	"fmt"
	"math"
)

func main() {
	/* Pembulatan */
	fmt.Println(math.Round(1.75))
	fmt.Println(math.Floor(1.75))
	fmt.Println(math.Ceil(1.75))

	fmt.Println(math.Round(1.1))
	fmt.Println(math.Floor(1.1))
	fmt.Println(math.Ceil(1.1))

	/* Min Max */
	fmt.Println(math.Min(10, 12))
	fmt.Println(math.Max(10, 12))

	/* Pangkat */
	fmt.Println(math.Pow(10, 3))
	fmt.Println(math.Pow(5, 25))
}
