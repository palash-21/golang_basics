/* we can also create interfaces without any methods, known as empty interfaces
Normally in functions, when we pass values to the function parameters, we need to specify the data type of parameters in a function definition.
However, with an empty interface, we can pass parameters of any data type
*/

package main

import "fmt"

// empty interface as function parameter
func displayValue(i interface{}) {
	fmt.Println(i)
}

func main() {

	a := "Welcome to Programiz"
	b := 20
	c := false

	// pass string to the function
	displayValue(a)

	// pass integer number to the function
	displayValue(b)

	// pass boolean value to the function
	displayValue(c)

	// passing multiple arguments to function
	displayMultiValue(a, b, c)
}

// function with an empty interface as its parameter
func displayMultiValue(i ...interface{}) {
	fmt.Println(i...)
}
