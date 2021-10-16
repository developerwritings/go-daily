package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
	time.Sleep(1 * time.Second)
}
