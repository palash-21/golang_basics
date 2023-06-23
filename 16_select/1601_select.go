/* The select statement in Go allows us to execute a channel among many alternatives
The syntax of select case looks similar to that of the Switch Case in Go.
And, like the switch case, only one of the cases is executed by select.
*/

package main

import "fmt"

func main() {

	// create two channels
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
	}

}

// goroutine to send integer data to channel
func channelNumber(number chan int) {
	number <- 15
}

// goroutine to send string data to channel
func channelMessage(message chan string) {
	message <- "Learning Go select"
}

// When you run this program, you might get different outputs.
// In this example, both channels are ready for execution, so the select statement executes the channel randomly.
