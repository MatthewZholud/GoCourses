package main

import (
	"fmt"
	"sort"
)

func printSorted(m map[int]string) {
	keys := make([]int, 0, len(m))
	for a := range m {
		keys = append(keys, a)
	}
	sort.Ints(keys)
	for _, a := range keys {
		fmt.Println(m[a])
	}
}

func main() {
	//m := map[int]string{2: "a", 0: "b", 1: "c"}
	m := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	printSorted(m)
}
