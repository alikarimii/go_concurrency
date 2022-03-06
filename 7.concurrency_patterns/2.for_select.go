package concurrencypatterns

import "fmt"

func ForSelctSample() {
	done := make(chan struct{})
	stringStream := make(chan string)
	// Sending iteration variables out on a channel
	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			return
		case stringStream <- s:
		}
	}

	// Looping infinitely waiting to be stopped
	for {
		// When we enter the select statement,
		// if the done channel hasn’t been closed,
		//  we’ll execute the default clause instead.
		select {
		case <-done:
			return
		default:
			fmt.Print("Do non-preemptable work here")
		}
		// or Do non-preemptable work here
	}
}
