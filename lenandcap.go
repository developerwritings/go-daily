// This program relates difference between len and capacity

package main

import "fmt"

func main() {
	s := make([]int, 0, 2)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("cap %v, len %v, %p\n", len(s), s)
	}
}
