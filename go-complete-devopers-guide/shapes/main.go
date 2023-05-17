package main

import "fmt"

type shape interface {
	area() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func main() {
	printArea(triangle{base: 5, height: 5})
	printArea(square{sideLength: 100})
}
