package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`v([a-z])l`)
	fmt.Println(regex.MatchString("nal"))
	fmt.Println(regex.MatchString("vel"))
	fmt.Println(regex.MatchString("xyz"))
	fmt.Println(regex.FindAllString("nal val nel vel vol xyz", 10))

	regex = regexp.MustCompile(`^[a-z]+$`)
	fmt.Println(regex.MatchString("noval"))
	fmt.Println(regex.MatchString("12noval"))
	fmt.Println(regex.MatchString("wardhana19"))
	fmt.Println(regex.MatchString("novalwardhana"))
	fmt.Println(regex.FindAllString("noval", 11))
}
