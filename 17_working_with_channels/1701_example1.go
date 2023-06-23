/* We can use channels to synchronize execution across goroutines.
Here’s an example of using a blocking receive to wait for a goroutine to finish.
*/

package main

import (
	"fmt"
	"time"
)

// The done channel will be used to notify another goroutine that this function’s work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true // Send a value to notify that we’re done.
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done

	/* If you removed the <- done line from this program,
	the program would exit before the worker even started. */
}

// When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup
