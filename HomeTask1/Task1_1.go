package main

import (
	"fmt"
	"sort"
)

func median(i []int) float64 {
	sort.Ints(i)
	Num := len(i) / 2
	if Odd(len(i)) {
		return float64(i[Num])
	}
	return float64(i[Num-1] + i[Num]) / 2
}

func Odd(number int) bool {
	return number%2 != 0
}

func main() {
	sl := []int{156, 2, 7, 3, 56, 0}
	s2 := []int{11, 12, 2, 1, -23}
	fmt.Printf("Median of s1: %f\nMedian of s2: %f", median(sl), median(s2))
}