// Slice is like an array but the length of a slice can grow and shrink as you see fit

package main

import (
	"fmt"
)

func main() {

	// initialzing a slice
	myslice1 := []int{} // an empty slice of 0 length and 0 capacity
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice := []int{1, 2, 3} // slice of length 3 and capacity also 3
	fmt.Println(myslice)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// creating slice from array
	var myarray = [5]int{10, 20, 30, 40, 50} // An array
	myslice3 := myarray[1:3]                 // A slice made from the array
	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3)) // capacity is 1 to the length of the array and slice can extend upto it

	myslice4 := myarray[0:3]
	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	// creating slice using make function
	// slice_name := make([]type, length, capacity)    --> capacity is optional, if not provided then will be equal to length
	myslice5 := make([]int, 5, 10)
	fmt.Printf("myslice5 = %v\n", myslice5)
	fmt.Printf("length = %d\n", len(myslice5))
	fmt.Printf("capacity = %d\n", cap(myslice5))

	myslice6 := make([]int, 5)
	fmt.Printf("myslice6 = %v\n", myslice6)
	fmt.Printf("length = %d\n", len(myslice6))
	fmt.Printf("capacity = %d\n", cap(myslice6))

	// modify slice
	myslice7 := []int{1, 3, 3, 4, 5, 6}
	fmt.Printf("myslice7 = %v\n", myslice7)
	fmt.Printf("length = %d\n", len(myslice7))
	fmt.Printf("capacity = %d\n", cap(myslice7))

	myslice7[1] = 2

	myslice7 = append(myslice7, 7, 8)
	fmt.Printf("myslice7 = %v\n", myslice7)
	fmt.Printf("length = %d\n", len(myslice7))
	fmt.Printf("capacity = %d\n", cap(myslice7))

	myslice8 := append(myslice7, 9, 10)
	fmt.Printf("myslice8 = %v\n", myslice8)
	fmt.Printf("length = %d\n", len(myslice8))
	fmt.Printf("capacity = %d\n", cap(myslice8))

	// Append one slice to another
	// slice3 = append(slice1, slice2...)  --> here ... after 2nd slice is necessary
	// can only append slices with same data type
	myslice9 := append(myslice7, myslice3...)

	fmt.Printf("myslice3=%v\n", myslice9)
	fmt.Printf("length=%d\n", len(myslice9))
	fmt.Printf("capacity=%d\n", cap(myslice9))

	// copy function copies from src to dest. It returns the number of elements copied.
	// copy(dest, src)
	// When using slices, Go loads all the underlying elements into the memory.
	// Using copy with only required elements will reduce the memory usage

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	fmt.Printf("neededNumbers = %v\n", neededNumbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	numbersCopy := make([]int, len(neededNumbers))
	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	copy(numbersCopy, neededNumbers)
	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy)) //capacity and length is reduced

}
