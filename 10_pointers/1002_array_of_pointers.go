/* Here we declares ptr as an array of MAX integer pointers.
Thus, each element in ptr, now holds a pointer to an int value.
The following example makes use of three integers, which will be stored in an array of pointers
*/

package main

import "fmt"

const MAX int = 3

func main() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* assign the address of integer. */
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptr[i])
	}
}
