/* Channels in Go act as a medium for goroutines to communicate with each other.
We know that goroutines are used to create concurrent programs.
Concurrent programs can run multiple processes at the same time.
However, sometimes there might be situations where two or more goroutines need to communicate with one another.
In such situations, we use channels that allow goroutines to communicate and share resources with each other.
*/

package main

import "fmt"

func main() {

	// create channel using make function
	number := make(chan int)

	// access type and value of channel
	fmt.Printf("Channel Type: %T\n", number)
	fmt.Printf("Channel Value: %v", number) // the value of a channel is a memory address, which acts as a medium through which goroutines send and receive data to communicate.

	message := make(chan string)

	// function call with goroutine
	go channelData(number, message)

	// retrieve channel data
	fmt.Println("Channel Data:", <-number)
	fmt.Println("Channel Data:", <-message)

}

func channelData(number chan int, message chan string) {

	// send data into channel
	number <- 15
	message <- "Learning Go channel"

}
