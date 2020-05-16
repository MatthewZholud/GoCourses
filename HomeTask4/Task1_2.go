package main

import (
	"errors"
	"fmt"
)

var ErrInputSmall = errors.New("bad input. too small")

var ErrInputBig = errors.New("bad input. too big")

type Point struct {
	x, y int64
}

type Square struct {
	start Point
	a     int64
}

func End(x int64, y int64, a int64) (int64, int64, error) {
	if a < 0 || x < 0 || y < 0 {
		return 0, 0, ErrInputSmall
	}

	if a > 100 || x > 100 || y > 100 {
		return 0, 0, ErrInputBig
	}
	return x + a, y - a, nil
}

func Perimeter(a int64) (int64, error) {
	if a < 0 {
		return 0, ErrInputSmall
	}

	if a > 100 {
		return 0, ErrInputBig
	}
	return a * 4, nil
}

func Area(a int64) (int64, error) {
	if a < 0 {
		return 0, ErrInputSmall
	}

	if a > 100 {
		return 0, ErrInputBig
	}
	return a * a, nil
}

func main() {
	s := Square{Point{1, 1}, 5}
	FinalX, FinalY, _ := End(s.start.x, s.start.y, s.a)
	fmt.Printf("End point: (%v,%v)\n", FinalX, FinalY)
	p, _ := Perimeter(s.a)
	fmt.Printf("Perimeter: %v\n", p)
	a, _ := Area(s.a)
	fmt.Printf("Area: %v\n", a)
}
