// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	links := []string{
// 		"https://google.com",
// 		"https://facebook.com",
// 		"https://stackoverflow.com",
// 		"https://amazon.com",
// 	}

// 	c := make(chan string)

// 	for _, link := range links {
// 		go checkLink(link, c)
// 	}
// 	// fmt.Println(<-c)
// 	// fmt.Println(<-c)
// 	// fmt.Println(<-c)
// 	// fmt.Println(<-c)
// 	// fmt.Println(<-c)
// 	// fmt.Println(<-c)

// 	// for i := 0; i < len(links); i++ {
// 	// 	fmt.Println(<-c)
// 	// }

// 	for l := range c {
// 		go func(link string) {
// 			time.Sleep(5 * time.Second)
// 			checkLink(link, c)
// 		}(l)
// 	}
// }

// func checkLink(link string, c chan string) {
// 	time.Sleep(5 * time.Second)
// 	_, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println("Network may be down")
// 		c <- "Network may be down"
// 		return
// 	}
// 	fmt.Println(link, ": Its up")
// 	c <- link
// }

package main

import "fmt"

func main() {
	c := make(chan string)
	for i := 0; i < 4; i++ {
		go printString("Hello there!", c)
	}

	for s := range c {
		fmt.Println(s)
	}
}

func printString(s string, c chan string) {
	fmt.Println(s)
	c <- "Done printing."
}
