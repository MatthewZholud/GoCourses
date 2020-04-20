package main

import "fmt"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() (int, int) {
	return s.start.x + int(s.a), s.start.y - int(s.a)
}

func (s Square) Perimeter() uint {
	return s.a * 4
}

func (s Square) Area() uint {
	return s.a * s.a
}

func displayPair(a, b interface{}) string {
	return fmt.Sprint(a, b)
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Printf("End point: (%v)\n", displayPair(s.End()))
	fmt.Printf("Perimeter: %v\n", s.Perimeter())
	fmt.Printf("Area: %v\n", s.Area())
}
