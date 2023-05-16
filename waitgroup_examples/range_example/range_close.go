package main

import "fmt"

func main() {
	c := make(chan int, 10)
	go test(cap(c), c)
	for newVal := range c {
		fmt.Println("hello", newVal)
	}

}
func test(val int, c chan int) {

	for i := 0; i < val; i++ {
		c <- i
	}
	close(c)
}
