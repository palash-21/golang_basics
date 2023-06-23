/* We have a function which has a long running code that is called from main.
We must terminate the long running function when a cancellation signal is sent from the caller
*/

/* Below is the commented program for normal function execution without cancellation.
So the func longRunning will run and complete its execution even if its not needed and can be cancelled at a point of time.

package main

import (
	"fmt"
	"time"
)

func longRunning() int {
	count := 0
	for i := 0; i < 5; i++ {
		count = count + i
		fmt.Println("Current value of count:", count)
		time.Sleep(2 * time.Second)

	}
	return count
}

func main() {
	count := longRunning()
	fmt.Println("count is ", count)
}
*/

// The first step towards the above goal is to create a new context in the main function.
// We should also be able to cancel this context after 5 seconds so that the longRunning will return.

// A context with a cancellation signal can be created using the WithCancel function.
// This function needs a context to be passed as a parameter.
// This is where the background context has its place.
// The background is a non-nil empty context.
// It is usually used as the starting point to create any new context.

package main

import (
	"context"
	"fmt"
	"time"
)

// passing the context from main as arguments to our function
// The context has a done channel which is used to notify context cancellation.
// The context's done channel which will be closed when the cancelFunc is called from main.
// This can be used in the longRunning function and it will know when the context is cancelled.
func longRunning(ctx context.Context) (int, error) {
	count := 0
	for i := 0; i < 5; i++ {
		select {
		// we check whether the done channel is closed
		case <-ctx.Done():
			// The reason why the done channel was closed can be found out by calling the Err() method on the context. And the error is returned.
			return 0, ctx.Err()
		default:
			count = count + i
			fmt.Println("Current value of count:", count)
			time.Sleep(2 * time.Second)
		}
	}
	return count, nil
}

func main() {
	// The WithCancel function returns a context and a CancelFunc.
	ctx, cancelFunc := context.WithCancel(context.Background())

	//  We create a new goroutine which will call the cancelFunc after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)

		// Calling the CancelFunc will send the cancellation signal.
		cancelFunc()
	}()
	count, err := longRunning(ctx)
	if err != nil {
		fmt.Println("long running task exited with error", err)
		return
	}

	fmt.Println("count is ", count)
}

/* We also can achieve the above objective by using WithTimeout function of the context package like :
ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
defer cancelFunc()


This will also cancel the function after 2 seconds.
*/
