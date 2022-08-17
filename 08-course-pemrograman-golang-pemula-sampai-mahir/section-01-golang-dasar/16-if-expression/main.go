package main

import "fmt"

func main() {
	/* If expression */
	var name = "Noval"
	if name == "noval" {
		fmt.Println(`my name is noval`)
	}
	if name == "Noval" {
		fmt.Println("my name is Noval")
	}

	/* If-else expression */
	var club = "chelsea"
	if club == "chelsea" {
		fmt.Println("Wow, chelsea")
	} else {
		fmt.Println("Other club")
	}

	/* If-else-if-else expression */
	var food = "soto"
	if food == "sate" {
		fmt.Println("Ini sate")
	} else if food == "bubur ayam" {
		fmt.Println("Ini bubur ayam")
	} else if food == "soto" {
		fmt.Println("Ini soto")
	} else {
		fmt.Println("Food not found")
	}

	/* If short statement: deklarasi variabel di dalam if */
	var fullname = "Novalita Kusuma Wardhana"
	if lengthName := len(fullname); lengthName > 20 {
		fmt.Println("Name too long")
	} else {
		fmt.Println("Name standard")
	}

}
