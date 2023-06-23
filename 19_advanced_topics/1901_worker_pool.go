/* In this example we’ll look at how to implement a worker pool using goroutines and channels. */

package main

import (
	"fmt"
	"time"
)

// Here’s the worker, of which we’ll run several concurrent instances.
// These workers will receive work on the jobs channel and send the corresponding results on results.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {

		// start the job
		fmt.Println("worker", id, "started  job", j)

		// wait for a second to stimulate work execution
		time.Sleep(time.Second)

		// finish and return result
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// 5 jobs each taking 1 sec
	const numJobs = 5

	// In order to use our pool of workers we need to send them work and collect their results.
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 jobs and then close that channel to indicate that’s all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work. This also ensures that the worker goroutines have finished.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

// Our running program shows the 5 jobs being executed by various workers.
// The program only takes about 2 seconds despite doing about 5 seconds of total work because there are 3 workers operating concurrently.
