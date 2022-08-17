package main

import (
	"container/ring"
	"fmt"
	"strconv"
	"time"
)

func main() {
	/*
		Container Ring
		1. Implementasi dari data circular list/lingkaran
		2. Adalah struktur data ring, dimana diakhir element akan kembali ke element awal
	*/
	var data *ring.Ring = ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data ke - " + strconv.Itoa(i)
		data = data.Next()
	}
	IterationFunc := func(n interface{}) {
		fmt.Println(n)
	}
	data.Do(IterationFunc)

	var data2 *ring.Ring = ring.New(10)
	for i := 0; i < data2.Len(); i++ {
		data2.Value = "Time: " + time.Now().Format("2006-01-02 15:04:05")
		data2 = data2.Next()
	}
	data2.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
