/* While running multiple Goroutines together, they interact with each other concurrently.
The order of execution of goroutine is random, so we might not get the output as expected
*/

// Program to illustrate multiple goroutines

package main

import (
	"fmt"
	"time"
)

func display(message string) {

	fmt.Println(message)

}

func main() {

	// run two different goroutine
	go display("Process 1")
	go display("Process 2")
	go display("Process 3")

	// to sleep main goroutine for 1 sec
	time.Sleep(time.Second * 1)
}
