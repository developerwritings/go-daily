package main

import (
	"container/list"
	"log"
)

type CustomStack struct {
	stack list.List
}

func (cs *CustomStack) push(element string) {
	cs.stack.PushFront(element)
}

func (cs *CustomStack) pop() {
	frontElement := cs.stack.Front()
	log.Println(frontElement)
	cs.stack.Remove(frontElement)
}

func main() {
	cs := &CustomStack{
		stack: *list.New(),
	}
	cs.push("A")
	cs.push("B")
	log.Println(cs.stack)
	for cs.stack.Len() > 0 {
		cs.pop()
	}
	log.Println(cs.stack)
}
