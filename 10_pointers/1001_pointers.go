/* Some Go programming tasks are performed more easily with pointers, and other tasks, such as call by reference, cannot be performed without using pointers.
Every variable is a memory location and every memory location has its address defined which can be accessed using ampersand (&) operator, which denotes an address in memory.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int = 10
	fmt.Printf("Address of a variable: %x\n", &a)

	//A  pointer is a variable whose value is the address of another variable, i.e., direct address of the memory location.
	// Like any variable or constant, you must declare a pointer before you can use it to store any variable address.
	fmt.Println(strings.Repeat("-", 10))
	// The general form of a pointer variable declaration is :−  var var_name *var-type
	// var-name is the name of the pointer variable and var-type is the datatype of the variable which the pointer points to
	// the actual datatype of pointer is fixed : a long hexadecimal number that represents a memory address
	var b int = 20 /* actual variable declaration */
	var ip *int    /* pointer variable declaration */

	ip = &b /* store address of b in pointer variable*/

	fmt.Printf("Address of b variable: %x\n", &b)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	// changing variable value using pointer
	*ip = 2
	fmt.Printf("Value b variable: %x\n", &b)

	fmt.Println(strings.Repeat("-", 10))
	/* While passing pointers to a function, we are actually passing a reference (address) of the variable.
	Instead of working with the actual value, we are working with references like
			-accessing value using reference
			-changing value using reference
	That's why this process of calling a function with pointers is called call by reference in Go
	*/
	var number int

	// passing value to function defined below
	callByValue(number)

	// passing a reference (address) to function defined below
	callByReference(&number)

	// Go compiler assign a Nil value to a pointer variable in case you do not have exact address to be assigned.
	// This is done at the time of variable declaration.
	// A pointer that is assigned nil is called a nil pointer
	fmt.Println(strings.Repeat("-", 10))
	var ptr *int
	fmt.Printf("The value of ptr is : %x\n", ptr)

	/* On most of the operating systems,
	programs are not permitted to access memory at address 0 because that memory is reserved by the operating system.
	However, the memory address 0 has special significance;
	it signals that the pointer is not intended to point to an accessible memory location.
	But by convention, if a pointer contains the nil (zero) value, it is assumed to point to nothing.
	*/
}

// call by value
func callByValue(num int) {

	num = 30
	fmt.Println(num) // 30

}

// call by reference
func callByReference(num *int) {

	*num = 10
	fmt.Println(*num) // 10

}
