package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	ch1 := make(chan int)

	go test(1000, ch1)
	for val := range ch1 {
		fmt.Println(val)
	}

}

func test(val int, ch chan int) {
	for i := 1; i <= val; i++ {
		if i%2 == 0 {
			ch <- i
		} else {

		}
	}
	close(ch)
}
