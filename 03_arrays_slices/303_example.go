/* A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.
*/

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

/* A slice literal is like an array literal without the length
When we create a slice directly, it creates an array then builds a slice that references it.

The zero value of a slice is nil.
A nil slice has a length and capacity of 0 and has no underlying array.

When appending elements to array,
If the backing array of slice is too small to fit all the given values a bigger array will be allocated.
The returned slice will point to the newly allocated array.
*/
