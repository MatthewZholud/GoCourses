package main

import (
	"errors"
	"fmt"
	"math"
)

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return s.side * 4
}

type Circle struct {
	rad float64
}

func (s Circle) area() float64 {
	return s.rad * s.rad * math.Pi
}

func (s Circle) perimeter() float64 {
	return 2 * s.rad * math.Pi
}

func InputVal(val float64) float64 {
	if val < 0 {
		ErrNegVal := errors.New("Figure elements cant have negative values")
		panic(ErrNegVal)
	}
	return val
}

func main() {
	var s Figure = Square{InputVal(5.5)}
	var c Figure = Circle{InputVal(3)}

	fmt.Printf("Area: %v Perimeter: %v\n", s.area(), s.perimeter())
	fmt.Printf("Area: %.5v Perimeter: %.5v\n", c.area(), c.perimeter())

}
