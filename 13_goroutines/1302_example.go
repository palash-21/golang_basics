/* In a concurrent program, the main() is always a default goroutine.
Other goroutines can not execute if the main() is not executing.
So, in order to make sure that all the goroutines are executed before the main function ends,
we sleep the process so that the other processes get a chance to execute.
*/

// Program to run two goroutines concurrently

package main

import (
	"fmt"
	"time"
)

// create a function
func display(message string) {

	// for loop
	for i := 0; i < 3; i++ {
		fmt.Println(message)

		// to sleep the process for 1 second
		time.Sleep(time.Second * 1)
	}
}

func main() {

	// call function with goroutine
	go display("Process 1")

	display("Process 2")

}
