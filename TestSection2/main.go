package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func printArea(s shape) {
	fmt.Println("Area is ", s.getArea())
}

func main() {
	sq := square{sideLength: 4}
	tr := triangle{
		height: 2,
		base:   2,
	}
	printArea(sq)
	printArea(tr)
}
