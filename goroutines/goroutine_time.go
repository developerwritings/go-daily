package main

import (
	"fmt"
	"time"
)

func main() {
	go routine1()
	go routine2()
	// we are stoppping one second to complete go Routine execution
	time.Sleep(1 * time.Second)
}

func routine1() {
	fmt.Println("Routine 1 ")
}

func routine2() {
	fmt.Println("Routine 1 ")
}
