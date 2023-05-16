package main

import "fmt"

type Base interface {
	hello()
}

type news struct {
	Val int
	Base
}

func (c *news) hello() {
	fmt.Println("fdfdfdf")
}

// func hello(c Sample) {
// 	fmt.Println(c.Val)
// }

func main() {
	news := &news{
		Val: 10,
	}
	news.hello()
}
