package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan bool)
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		ch <- true
	})

	for time.Since(start) < 5*time.Second {
		<-ch
		t.Reset(randomDuration())
	}
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
