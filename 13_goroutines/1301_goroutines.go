/* A goroutine is a lightweight thread managed by the Go runtime.
Goroutines run in the same address space, so access to shared memory must be synchronized.

We use goroutines to create concurrent programs.
Concurrent programs are able to run multiple processes at the same time.

We can convert a regular function to a goroutine by calling the function using the go keyword
*/

package main

import "fmt"

// create a function
func display(message string) {

	fmt.Println(message)
}

func main() {

	// call goroutine
	go display("Process 1")

	display("Process 2")
}

/* In our example, the next statement is the second call to the function. So, first Process 1 should be printed and then Process 2.
However, we are only getting Process 2 as output.
This is because we have used go with the first function call, so it is treated as a goroutine.
And the function runs independently and the main() function now runs concurrently.
Hence, the second call is executed immediately and the program terminates without completing the first function call.
*/
