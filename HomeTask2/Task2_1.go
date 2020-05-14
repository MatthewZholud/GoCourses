package main

import "fmt"

func average(arr [6]int) float64 {
	var sum float64
	for _, a := range arr {
		sum += float64(a)
	}
	return sum / float64(len(arr))
}

func main() {
	var arr = [6]int{4, -7, 3, -67, 2, 13}
	fmt.Printf("Average value of array: %v", average(arr))
}
