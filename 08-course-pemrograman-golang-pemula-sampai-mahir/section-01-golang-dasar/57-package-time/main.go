package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Time
		Package yang berisi fungsionalitas untuk manajemen waktu di golang
	*/

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond())

	nowDate := time.Date(2021, time.November, 28, 17, 30, 59, 100, time.UTC)
	fmt.Println(nowDate)

	timeParse, err := time.Parse("2006-01-02 15:04:05", "2021-11-28 17:30:59")
	if err != nil {
		fmt.Println("Error time parse: ", err)
	} else {
		fmt.Println("Success time parse: ", timeParse)
	}
	timeParse, err = time.Parse("02/01/2006 15:04:05", "32/15/2022 17:30:59")
	if err != nil {
		fmt.Println("Error time parse: ", err)
	} else {
		fmt.Println("Success time parse: ", timeParse)
	}
	timeParse, err = time.Parse("02/01/2006 15:04:05", "29/11/1992 23:00:00")
	if err != nil {
		fmt.Println("Error time parse: ", err)
	} else {
		fmt.Println("Success time parse: ", timeParse)
	}
}
