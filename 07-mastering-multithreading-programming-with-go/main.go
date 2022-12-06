// You can edit this code!
// , here and start typing.
package main

import (
	"fmt"
	"sync"
)

func count() {
	wg := sync.WaitGroup{}
	x := 0
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go increment(&x, &wg)
	}
	wg.Wait()
	fmt.Printf("%d\n", x)
}

func increment(x *int, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		*x += 1
	}
	wg.Done()
}

func main() {
	count()
}
