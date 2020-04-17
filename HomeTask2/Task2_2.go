package main

import "fmt"

func max(sl []string) string {
	word := sl[0]
	for _, a := range sl {
		if len(a) > len(word) {
			word = a
		}
	}
	return word
}

func main() {
	sl := []string{"one", "two", "three"}
	//sl := []string{"one", "two"}
	fmt.Printf("The longest word: %v", max(sl))
}
