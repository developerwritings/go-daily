package main

import (
	"container/list"
	"log"
)

type CustomQueue struct {
	queue list.List
}

func (cq *CustomQueue) Enque(element string) {
	cq.queue.PushBack(element)
}

func (cq *CustomQueue) Deque() {
	frontElement := cq.queue.Front()
	cq.queue.Remove(frontElement)
}

func main() {
	cq := CustomQueue{
		queue: *list.New(),
	}

	cq.Enque("A")
	cq.Enque("B")
	cq.Enque("C")
	cq.Enque("D")
	log.Println(cq.queue)
	for cq.queue.Len() > 0 {
		cq.Deque()
	}
	log.Println(cq.queue)

}
