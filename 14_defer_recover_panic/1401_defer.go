/* We use the defer statement to prevent the execution of a function until all other functions execute.*/

package main

import "fmt"

func main() {

	// defer the execution of Println() function

	fmt.Println("One")
	defer fmt.Println("Three")
	defer fmt.Println("Two")
	// When we use multiple defer in a program, the order of execution of the defer statements will be LIFO (Last In First Out)
}
