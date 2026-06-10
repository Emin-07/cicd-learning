package main

import (
	"fmt"
	"sync"
)

var x int

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for ; x < 5; x++ {
		go func() {
			fmt.Println(x)
			wg.Done()
		}()
	}

	wg.Wait()
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
