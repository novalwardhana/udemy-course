package main

import "fmt"

func main() {
	/*
		Closure
		CLosure adalah kemampuan sebuah fungsi untuk berinteraksi untuk berinteraksi dengan data-data disekitarnya dalam scope yang sama
	*/

	var counter int
	var myName = "Noval"
	increament := func() {
		myName := "Kusuma"
		fmt.Println(myName)
		counter++
	}

	increament()
	increament()
	fmt.Println(myName)
	fmt.Println(counter)
}
