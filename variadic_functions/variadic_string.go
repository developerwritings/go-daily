// variadic function
// vardiac functions can be called with multiple argument of same type at once

package main

import "fmt"

func main() {
	vardiac_string_example("hello", "Good Morning")
}

func vardiac_string_example(str ...string) {
	fmt.Println(str)
}
