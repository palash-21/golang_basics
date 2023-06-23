/* The panic statement immediately terminates the program.
However, sometimes it might be important for a program to complete its execution and get some required results.
In such cases, we use the recover statement to handle panic in Go.
The recover statement prevents the termination of the program and recovers the program from panic.
*/

package main

import "fmt"

// recover function to handle panic
func handlePanic() {

	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}

}

func division(num1, num2 int) {

	// execute the handlePanic even after panic occurs
	defer handlePanic()

	// if num2 is 0, program is terminated due to panic
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
}

func main() {

	division(4, 2)
	division(8, 0)
	division(19, 8)

}

/* We are calling the handlePanic() before the occurrence of panic.
It's because the program will be terminated if it encounters panic and we want to execute handlePanic() before the termination.
We are using defer to call handlePanic().
It's because we only want to handle the panic after it occurs, so we are deferring its execution
*/
