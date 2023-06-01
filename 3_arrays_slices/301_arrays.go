// Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value

// length - fixed, type of all elements - same

package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3}    // array of length 3 and elements of type int
	arr2 := [5]int{4, 5, 6, 7, 8} // array of length 5 and elements of type int

	var arr3 = []int{1, 2, 3} // length is inferred from number of elements
	arr4 := []int{4, 5, 6, 7, 8}

	fmt.Println(arr1, arr2, arr3, arr4)

	var arr5 = []string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(arr5)
	arr5[3] = "Audi" // changing elements in array (index starts from 0)
	fmt.Println(arr5)

	// initialization
	arr6 := [5]int{}              //not initialized
	arr7 := [5]int{1, 2}          //partially initialized
	arr8 := [5]int{1, 2, 3, 4, 5} //fully initialized
	arr9 := [5]int{1: 2, 2: 3}    // initializing only specific elements

	fmt.Println(arr6, arr7, arr8, arr9)

	// len() to get length of array
	fmt.Println(len(arr9))
}
