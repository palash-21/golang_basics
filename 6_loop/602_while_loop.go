// Unlike other programming languages, Go doesn't have a dedicated keyword for a while loop.
// However, we can use the for loop to perform the functionality of a while loop.

// Program to create a multiplication table of 5 using while loop

package main

import (
	"fmt"
)

func main() {
	multiplier := 1

	// run while loop for 10 times
	for multiplier <= 10 {

		// find the product
		product := 5 * multiplier

		// print the multiplication table in format 5 * 1 = 5
		fmt.Printf("5 * %d = %d\n", multiplier, product)
		multiplier++
	}

	// Go do...while Loop
	number := 1

	// loop that runs infinitely
	for {

		// condition to terminate the loop
		if number > 5 {
			break
		}

		fmt.Printf("%d\n", number)
		number++

	}

}
