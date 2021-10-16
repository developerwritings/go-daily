package main

import "fmt"

func main() {
	channel := make(chan int)
	go channelTest(10, channel)
	val := <-channel
	fmt.Println(val)
}

func channelTest(val int, c chan int) {

	val = val * val
	c <- val
}
