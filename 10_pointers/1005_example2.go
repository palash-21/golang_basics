// Program to pass pointer as a function argument

package main

import "fmt"

// function definition with a pointer parameter
func update(num *int) {

	// dereference the pointer (change the value using pointer)
	*num = 30

}

func main() {

	var number = 55

	// function call
	// passing address of variable to update function
	update(&number)

	fmt.Println("The number is", number)

}
