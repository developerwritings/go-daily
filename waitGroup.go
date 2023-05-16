package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	go PrintNumbers(ch, 10, &wg)
	for val := range ch {
		fmt.Println(val)
	}
	defer wg.Wait()
}

func PrintNumbers(ch chan int, val int, wg *sync.WaitGroup) {
	wg.Add(1)
	for i := 0; i < val; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}
