package main

import "fmt"

func main() {

	/* For like while loop */
	counter := 1
	counterMax := 7
	for counter <= counterMax {
		fmt.Println("Iteration: ", counter)
		counter++
	}

	/* For with statement */
	for counter := 1; counter <= counterMax; counter++ {
		fmt.Println("Iteration: ", counter)
	}

	/* For range */
	clubs := make([]string, 10)
	clubs[0] = "Chelsea"
	clubs[1] = "Arsenal"
	clubs[2] = "Liverpool"
	clubs[5] = "Everton"
	for index, club := range clubs {
		fmt.Println(index, ": ", club)
	}
	person := make(map[string]string)
	person["name"] = "Noval"
	person["hoby"] = "Programming"
	person["age"] = "28"
	for key, data := range person {
		fmt.Println(key, ": ", data)
	}
	fmt.Println(person)
}
