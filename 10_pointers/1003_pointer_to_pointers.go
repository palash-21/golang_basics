/* A pointer to a pointer is a form of chain of pointers.
Normally, a pointer contains the address of a variable.
When we define a pointer to a pointer, the first pointer contains the address of the second pointer, which points to the location that contains the actual value

Pointer to Pointer(Address)     --->     Pointer(Address)    --->   Variable(value)
*/

package main

import "fmt"

func main() {
	var a int      // variable a
	var ptr *int   // pointer to variable a
	var pptr **int // pointer to the pointer

	a = 3000

	/* take the address of var */
	ptr = &a

	/* take the address of ptr using address of operator & */
	pptr = &ptr

	/* take the value using pptr */
	fmt.Printf("Value of a = %d\n", a)
	fmt.Printf("Address available at ptr = %d\n", ptr)
	fmt.Printf("Value available at *ptr = %d\n", *ptr)
	fmt.Printf("Address available at *pptr = %d\n", *pptr)
	fmt.Printf("Value available at **pptr = %d\n", **pptr)

}
