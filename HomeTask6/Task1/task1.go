package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)
	var Lock sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			Lock.Lock()
			defer Lock.Unlock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
