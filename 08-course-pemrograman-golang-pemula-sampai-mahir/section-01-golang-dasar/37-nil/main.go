package main

import "fmt"

func generateMap(param string) map[string]string {
	if param == "" {
		return nil
	} else {
		return map[string]string{}
	}
}

func generateSlice(param string) []int {
	if param == "" {
		return nil
	} else {
		return []int{1, 4, 9, 16, 25}
	}
}

func generateInterface(param string) interface{} {
	if param == "" {
		return nil
	} else {
		type Person struct {
			FirstName string
			LastName  string
			Address   int
		}
		return Person{
			FirstName: "Noval",
			LastName:  "Wardhana",
			Address:   28,
		}
	}
}

func main() {
	/*
		Nil
		1. Biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
		2. Berbeda dengan golang, di golang saat kita buat variabel dengan tipe data tertentu maka secara otomatis kana dibuatkan default valuenya
		3. Di golang ada data nil yaitu data kosong
		4. Nil sendiri bisa digunakan di beberapa tipe data seperti interface, function, map, slice, pointer, dan channel
	*/

	/* Nil Map */
	newMap := generateMap("")
	if newMap == nil {
		fmt.Println("Nil map")
	} else {
		fmt.Println("New Map")
	}
	newMap = generateMap("noval")
	if newMap == nil {
		fmt.Println("Nil map")
	} else {
		fmt.Println("New Map")
	}

	/* Nil slice */
	newSlice := generateSlice("test")
	if newSlice == nil {
		fmt.Println("Nil slice")
	} else {
		fmt.Println("New slice: ", newSlice)
	}
	newSlice = generateSlice("")
	if newSlice == nil {
		fmt.Println("Nil slice")
	} else {
		fmt.Println("New slice: ", newSlice)
	}

	/* Nil interface */
	newInterface := generateInterface("")
	if newInterface == nil {
		fmt.Println("Nil interface")
	} else {
		fmt.Println("New interface: ", newInterface)
	}
	newInterface = generateInterface("test")
	if newInterface == nil {
		fmt.Println("Nil interface")
	} else {
		fmt.Println("New interface: ", newInterface)
	}
}
