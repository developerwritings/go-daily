package main

import (
	"log"
	"sync"
)

type CustomSliceStack struct {
	stack []string
	lock  sync.RWMutex
}

func (css *CustomSliceStack) push(element string) {
	css.stack = append(css.stack, element)

}

func (css *CustomSliceStack) pop() {
	css.stack = css.stack[:len(css.stack)-1]
	log.Println(css.stack)
}

func main() {
	css := &CustomSliceStack{
		stack: make([]string, 0),
	}

	css.push("A")
	css.push("B")
	css.push("C")
	log.Println(css.stack)
	for len(css.stack) > 0 {
		css.pop()
	}
	log.Println(css.stack)
}
