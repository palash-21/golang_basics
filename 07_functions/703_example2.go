package main

import (
	"fmt"
	"strings"
)

// function closures : like a function that returns another function

// Three examples :

// function for example 1
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	// example 1
	/* nextNumber is now a function with i as 0 */
	nextNumber := getSequence()

	/* invoke nextNumber to increase i by 1 and return the same */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* create a new sequence and see the result, i is 0 again*/
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	fmt.Println(strings.Repeat("-", 10))
	// example 2
	fmt.Println("example 2")

	// here, we will use adder function defined below
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fmt.Println(strings.Repeat("-", 10))
	// example 3
	fmt.Println("example 3")
	// print elements of fibonacci sequence
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// function for example 2
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

/* A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
For example, the adder function returns a closure. Each closure is bound to its own sum variable.
So here, closure returned to pos() is bound to a diff sum variable as compared to sum variable for neg()
*/

// function for example 3
// print elements of fibonacci sequence
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		r := x
		x, y = y, x+y

		return r
	}
}
