package main

import "fmt"

func main() {
	var result = 10 + 12 - 5
	fmt.Println(result)

	var a = 15
	var b = 13
	c := a + b
	fmt.Println(c)

	var j = 30
	j += 5
	j -= 3
	j *= 2
	j++
	fmt.Println(j)

	var k = true
	fmt.Println(!k)
	var l = false
	fmt.Println(!l)
}
