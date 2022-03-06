package concurrencypatterns_test

import (
	"testing"

	cp "github.com/alikarimii/go_concurrency/7.concurrency_patterns"
)

func TestBridge(t *testing.T) {

	for v := range cp.Bridge(nil, genVals()) {
		t.Logf("%v ", v)
	}
}

// Here’s an example that creates a series of 10 channels,
// each with one element written to them,
// and passes these channels into the bridge function
func genVals() <-chan <-chan interface{} {
	chanStream := make(chan (<-chan interface{}))
	go func() {
		defer close(chanStream)
		for i := 0; i < 10; i++ {
			stream := make(chan interface{}, 1)
			stream <- i
			close(stream)
			chanStream <- stream
		}
	}()
	return chanStream
}
