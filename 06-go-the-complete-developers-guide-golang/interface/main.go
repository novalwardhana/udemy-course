package main

import "fmt"

/*
	Did you run into any error messages as you worked on the program?
	If so, write the error message out here and a short explanation on what you think it might mean.

	yes, because we cannot add a reference in the interface or it will appear error "invalid receiver shape (pointer or interface type)".
	So the solution is to make shape interface a parameter inside printArea function

	I know how to make multiple struct fulfil into same interface

	Good, same with my solution that to make shape interface a parameter inside printArea function
*/

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	newSquare := square{
		sideLength: 10,
	}
	printArea(&newSquare)

	newTriangle := &triangle{
		height: 10,
		base:   5,
	}
	printArea(newTriangle)
}

func (s *square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t *triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
