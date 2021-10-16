package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Silent"
	str2 := "lisTen"

	if len(str) == len(str2) {
		for _, val := range str {
			val1 := strings.ToLower(string(val))
			if match(val1, str2) {
				result = "matching"
			} else {
				result = "not matching"
			}
		}
		fmt.Println(result)
	} else {
		fmt.Println("two string are not anagram")
	}
}

func match(charmatch string, stringMatch string) bool {

	var result bool = false
	for _, val := range stringMatch {
		val4 := strings.ToLower(string(val))
		if charmatch == val4 {
			result = true
		}
	}
	return result
}
