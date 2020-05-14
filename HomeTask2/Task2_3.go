package main

import (
	"fmt"
)

func reverse(slc []int64) []int64 {
	slc1 := make([]int64, cap(slc))
	copy(slc1, slc)
	for i, j := 0, len(slc1)-1; i < j; i, j = i+1, j-1 {
		slc1[i], slc1[j] = slc1[j], slc1[i]
	}
	return slc1
}

func main() {
	slc := []int64{1, 2, 12, 4, 5, 6, 93}
	fmt.Printf("Original: %v\nReversed: %v", slc, reverse(slc))

}
