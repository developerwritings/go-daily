package main

import "fmt"

func main() {
	v := make(chan int)
	// if v is nil, cap(v) is zero.
	// Channel: the channel buffer capacity, in units of elements;
	// if v is nil, cap(v) is zero.
	fmt.Println("Buffer channel capacity", cap(v))

	v1 := make(chan int, 10)

	fmt.Println("Buffer channel capacity", cap(v1))
}
