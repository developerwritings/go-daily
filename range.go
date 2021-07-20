package main

import "fmt"

func main() {

	// Range array of integers with indexes
	nums := []int{2, 3, 4}
	for i, v := range nums {

		fmt.Println("index value ", i, v)
	}

	// Range array of strings with indexes
	strs := []string{"developer", "writings"}
	for i, v := range strs {

		fmt.Println("index and values", i, v)
	}

	// Range over map values with indexes
	mapValues := map[string]string{"name": "developer", "value": "writings"}
	fmt.Println(mapValues)
	for i, v := range mapValues {
		fmt.Println("index values", i, v)
	}
}
