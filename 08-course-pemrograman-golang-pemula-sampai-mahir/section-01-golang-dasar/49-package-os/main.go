package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		Package OS
		1. Package OS berisi fungsionalitas untuk mengakses fitur sistem operasi secara independen
		2. Bisa untuk semua sistem operasi
	*/

	/* Argumen */
	args := os.Args
	fmt.Println(args)

	/* Hostname */
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Err Hostname: ", err)
	} else {
		fmt.Println("Hostname:", hostname)
	}

	/* Get ENV */
	goENV := os.Getenv("GO_ENV")
	fmt.Println(goENV)
	goPath := os.Getenv("GOPATH")
	fmt.Println(goPath)
	temp := os.Getenv("TEMP")
	fmt.Println(temp)
}
