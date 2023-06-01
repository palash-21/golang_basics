// Function to get factorial of a number

// Writing recursive function

package main

import (
	"fmt"
	"strings"
)

func get_factorial(num int) (result int) {
	if num == 0 {
		return 1
	} else {
		return num * get_factorial(num-1)
	}
}

func main() {
	fmt.Println(get_factorial(4))
	fmt.Println(get_factorial(6))
	fmt.Println(get_factorial(10))

	fmt.Println(strings.Repeat("-", 10))
	// Declaring a function fib
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
