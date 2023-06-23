/* By default channels are unbuffered,
meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
Buffered channels accept a limited number of values without a corresponding receiver for those values.
*/

package main

import "fmt"

func main() {

	// A channel of strings buffering up to 2 values.
	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	// Later we can receive these two values as usual.
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
