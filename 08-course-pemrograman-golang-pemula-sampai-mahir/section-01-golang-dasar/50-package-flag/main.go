package main

import (
	"flag"
	"fmt"
)

func main() {
	/*
		Package Flags
		Berisi fungsionalitas untuk mengakses command line argument
		go run . -name=noval -address=bantul
	*/
	var name *string = flag.String("name", "nov", "Put name data")
	var address *string = flag.String("address", "yk", "Put address data")
	var age *int = flag.Int("age", 29, "put age data")

	flag.Parse()

	fmt.Println("name:", *name)
	fmt.Println("address:", *address)
	fmt.Println("age:", *age)
}
