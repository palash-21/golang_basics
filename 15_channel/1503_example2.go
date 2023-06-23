/* When a goroutine receives data from a channel, the operation is blocked until another goroutine sends the data to the channel.
 */

package main

import "fmt"

func main() {

	// create channel
	ch := make(chan string)

	// function call with goroutine
	go receiveData(ch)

	fmt.Println("No data. Receive Operation Blocked")

	// send data to the channel
	ch <- "Data Received. Receive Operation Successful"

}

func receiveData(ch chan string) {

	// receive data from the channel
	fmt.Println(<-ch)

}

// When the channel sends data, the data received by goroutine is printed
