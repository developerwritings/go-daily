// variadic function
// vardiac functions can be called with multiple argument of same type at once
// It will help you caluclate the sum or arthametic operation
package main

import "fmt"

func main() {
	vardiac_int_example(1, 4, 10)
}

func vardiac_int_example(num ...int) {
	fmt.Println(num)
}
