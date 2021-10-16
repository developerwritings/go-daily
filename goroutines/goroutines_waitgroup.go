package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// Adding GoRoutine count to wait
		wg.Add(1)
		go func() {
			fmt.Println("go Routine started")
			wg.Done()
			// Decrement the counter once its completed goroutine job
		}()
	}
	wg.Wait()
	// it will all the go routines completes its execution
}
