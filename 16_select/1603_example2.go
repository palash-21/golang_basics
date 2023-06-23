/* We know that when both multiple channels are ready for execution, the select statement executes the channel randomly.
However, if only one of the channels is ready for execution, it executes that channel

The select statement blocks all channels if they are not ready for execution.
Suppose, in our previous example, if both the number and message channels are not ready for execution,
select blocks both the channels for a certain time until one is available for execution.

However, it's better to display some messages instead of blocking until the channels are ready.
For this, we use the default case, which is executed if none of the channels are ready for execution.
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	// create channels
	number := make(chan int)
	message := make(chan string)

	// function call with goroutine
	go channelNumber(number)
	go channelMessage(message)

	// selects and executes a channel
	select {
	case firstChannel := <-number:
		fmt.Println("Channel Data:", firstChannel)

	case secondChannel := <-message:
		fmt.Println("Channel Data:", secondChannel)

	// default case
	default:
		fmt.Println("Wait!! Channels are not ready for execution")
	}

}

// goroutine to send data to the channel
func channelNumber(number chan int) {

	// sleeps the process for 2 seconds
	time.Sleep(2 * time.Second)

	number <- 15
}

// goroutine to send data to the channel
func channelMessage(message chan string) {

	// sleeps the process by 2 seconds
	time.Sleep(2 * time.Second)

	message <- "Learning Go Select"
}
