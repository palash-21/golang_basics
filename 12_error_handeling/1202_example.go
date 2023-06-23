/* We can also handle Go errors using the Errorf() function. Unlike, New(), we can format the error message using Errorf().
This function is present inside the fmt package, so we can directly use this if we have imported the fmt package
*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	age := -14

	// create an error using Efforf()
	error := fmt.Errorf("%d is negative\nAge can't be negative", age)

	if age < 0 {
		fmt.Println(error)
	} else {
		fmt.Println("Age: %d", age)
	}

	// checking error for division by zero
	fmt.Println(strings.Repeat("*", 10))
	err := divide(4, 0)

	// error found
	if err != nil {
		fmt.Printf("error: %s", err)

		// error not found
	} else {
		fmt.Println("Valid Division")
	}
}

// Division by zero
func divide(num1, num2 int) error {

	// returns error
	if num2 == 0 {
		return fmt.Errorf("%d / %d\nCannot Divide a Number by zero", num1, num2)
	}

	// returns the result of division
	return nil
}
