package concurrencypatterns

func Bridge(
	done <-chan interface{},
	chanStream <-chan <-chan interface{},
) <-chan interface{} {
	// This is the channel that will return all values from bridge.
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		// This loop is responsible for pulling channels off
		// of chanStream and providing them to a nested loop for use.
		for {
			var stream <-chan interface{}
			select {
			case maybeStream, ok := <-chanStream:
				if !ok {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}
			// This loop is responsible for reading values off the channel
			// it has been given and repeating those values onto valStream.
			// When the stream we’re currently looping over is closed,
			// we break out of the loop performing the reads from this channel,
			// and continue with the next iteration of the loop,
			// selecting channels to read from. This provides us with an
			// unbroken stream of values.
			for val := range OrDone(done, stream) {
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()
	return valStream
}
