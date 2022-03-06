package concurrencypatterns

import "fmt"

func LexicalSample() {
	chanOwner := func() <-chan int {
		// Here we instantiate the channel within the lexical scope 
		// of the chanOwner function. This limits the scope of the write aspect
		// of the results channel to the closure defined below it. 
		// In other words, it confines the write aspect of this channel 
		// to prevent other goroutines from writing to it
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}
	// Here we receive a read-only copy of an int channel. 
	// By declaring that the only usage we require is read access, 
	// we confine usage of the channel within the con sume function to only reads
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}
	// Here we receive the read aspect of the channel and we’re able 
	// to pass it into the consumer, which can do nothing but read from it. 
	// Once again this confines the main goroutine to a read-only view of the channel.
	results := chanOwner()
	consumer(results)
}
