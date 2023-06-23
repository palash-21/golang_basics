// Program to return a pointer from a function

package main

import "fmt"

func main() {

	// function call
	result := display()

	// print value of variable using pointer
	fmt.Println("Welcome to", *result)

}

// function that returns pointer to a string variable
func display() *string {

	message := "Programiz"

	// returns the address of message
	return &message

}
