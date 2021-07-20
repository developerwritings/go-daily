package main

import (
	"fmt"
	"time"
)

func main() {
	hello()
	fmt.Println("Heloow")

}

func hello() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Hello")
	}
}
