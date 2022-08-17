package main

import (
	"fmt"
	"sort"
)

type Food struct {
	Name  string
	Price int
}

type FoodList []Food

func (f FoodList) Len() int {
	return len(f)
}

func (f FoodList) Less(i, j int) bool {
	return f[i].Price < f[j].Price
}

func (f FoodList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type Club struct {
	Name    string
	Country string
	WonUCL  int
}

type ClubList []Club

func (cl ClubList) Len() int {
	return len(cl)
}

func (cl ClubList) Less(i, j int) bool {
	return cl[i].WonUCL < cl[j].WonUCL
}

func (cl ClubList) Swap(i, j int) {
	cl[i], cl[j] = cl[j], cl[i]
}

func main() {
	/*
		Sort
		1. Package yang berisi utilitas untuk proses pengurutan
		2. Agar data bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort interface
	*/
	foods := FoodList{
		Food{"Sate", 25000},
		Food{"Bakso", 15000},
		Food{"Gurameh", 50000},
		Food{"Soto", 12000},
	}
	fmt.Println(foods)
	sort.Sort(foods)
	fmt.Println(foods)

	clubs := ClubList{
		Club{"Manchester United", "England", 3},
		Club{"Real Madrid", "England", 13},
		Club{"Chelsea", "England", 2},
		Club{"Juventus", "Italy", 2},
		Club{"Bayern Munich", "Germany", 5},
	}
	fmt.Println(clubs)
	sort.Sort(clubs)
	fmt.Println(clubs)
}
