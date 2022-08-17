package main

import "fmt"

func main() {
	/*
		Break: menghentikan seluruh perulangan
	*/
	for i := 0; i < 11; i++ {
		fmt.Println(i)
		if i == 9 {
			break
		}
	}

	/*
		Continue:langsung melanjutkan ke perulangan selanjutnya
	*/
	i := 0
	for {
		i++
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
		if i == 10 {
			break
		}
	}
}
