// Entry Point of the file we need to define package first default is main
package main

// Import the packages to use we are using standard pacakge fmt to Print Values
import "fmt"

// Main function is the entry point of the any Project
func main() {

	// Declare a pointer variable
	var s *int
	fmt.Println(s)

	/*
		Example 1
		Declared a variable with having some static value and modifying the value
	*/

	var i int = 4
	p := &i
	// Priting the value of p, It will print the address of the variable i
	fmt.Println(p)
	// Priting the value of pointer p, It will print the value of the variable i
	fmt.Println(*p)
	// Modifiying Pointer value
	*p = 25
	// Printing the value of the pointer
	fmt.Println(*p)
	// Verifying that address changes or it having same address
	fmt.Println(p)

	/*
		Example 2
		Declared a variable with out assiging values to variable
	*/

	var x int
	q := &x
	// Priting the value of p, It will print the address of the variable x
	fmt.Println(q)
	// Priting the value of pointer p, It will print the value of the variable x
	// It will Pring 0 because golang default assign value to zero
	fmt.Println(*q)
	// Modifiying Pointer value
	*q = 6
	// Printing the value of the pointer
	fmt.Println(*q)
	// Verifying that address changes or it having same address
	fmt.Println(q)

}
