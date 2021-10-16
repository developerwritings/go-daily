package main

import "fmt"

func main() {

	ch := make(chan int)
	go test(ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func test(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
