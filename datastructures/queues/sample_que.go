package main

import (
	"container/list"
	"fmt"
)

type customQueue struct {
	queue *list.List
}

func (c *customQueue) Size() int {
	return c.queue.Len()
}

func (c *customQueue) Enque(value string) {
	c.queue.PushBack(value)
}

func main() {
	customQueue := &customQueue{
		queue: list.New(),
	}
	customQueue.Enque("B")
	customQueue.Enque("C")

	fmt.Printf("Size: %d\n", customQueue.Size())

	// for customQueue.Size() > 0 {
	front_val := customQueue.queue.Front()
	fmt.Println(front_val)
	front_vals := customQueue.queue.Remove(front_val) // }
	fmt.Println(front_vals)
}
