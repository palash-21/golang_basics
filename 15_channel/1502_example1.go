/* In Go, the channel automatically blocks the send and receive operations depending on the status of goroutines.
When a goroutine sends data into a channel, the operation is blocked until the data is received by another goroutine.
*/

package main

import "fmt"

func main() {

	// create channel
	ch := make(chan string)

	// function call with goroutine
	go sendData(ch)

	// receive channel data
	fmt.Println(<-ch)

}

func sendData(ch chan string) {

	// data sent to the channel
	ch <- "Received. Send Operation Successful"
	fmt.Println("No receiver! Send Operation Blocked")

}

// When the channel is ready to receive data, the data sent by goroutine is printed.
