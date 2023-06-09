/* Go’s sort package implements sorting for builtins and user-defined types.
Lets first look at sorting for built-in types.
Sort methods are specific to the builtin type
*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	// Note that sorting is in-place, so it changes the given slice and doesn’t return a new one
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting ints.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use sort to check if a slice is already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
