package main

import (
	"log"
	"sync"
)

type CustomSliceQueue struct {
	queue []string
	lock  sync.RWMutex
}

func (csq *CustomSliceQueue) Enque(element string) {
	csq.queue = append(csq.queue, element)
}

func (csq *CustomSliceQueue) Deque() {
	csq.queue = csq.queue[1:]
}

func main() {
	csq := &CustomSliceQueue{
		queue: make([]string, 0),
	}
	csq.Enque("A")
	csq.Enque("B")
	log.Println(csq.queue)
	for len(csq.queue) > 0 {
		csq.Deque()
	}
	log.Println(csq.queue)
}
