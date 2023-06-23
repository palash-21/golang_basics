// go only has for loop

/* The for loop can take up to three statements:
statement1 :Initializes the loop counter value.
statement2 :Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends. (condition)
statement3 :Increases the loop counter value.
*/

package main

import (
	"fmt"
)

func main() {

	// for loop
	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}

	// for loop with continue and break
	for j := 0; j <= 10; j++ {
		if j == 3 {
			continue // jump to next iteration
		}
		if j == 7 {
			break // break the loop and exit
		}
		fmt.Println(j)
	}

	// The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value
	// Used like --> for index, value := range array|slice|map

	fruits := [3]string{"apple", "orange", "banana"}

	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	// goto

	var a int = 10
	/* do loop execution */
LOOP:
	for a < 20 {
		if a == 15 {
			/* skip the iteration */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}
