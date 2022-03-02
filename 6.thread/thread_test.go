package thread_test

import (
	"sync"
	"testing"
)

func BenchmarkThread(b *testing.B) {
	var wg sync.WaitGroup
	// it's used for
	begin := make(chan struct{})
	c := make(chan struct{})
	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin // Here we wait until we’re told to begin
		for i := 0; i < b.N; i++ {
			c <- token // Here we send messages to the receiver goroutine
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin // Here we wait until we’re told to begin
		for i := 0; i < b.N; i++ {
			<-c // Here we receive a message but do nothing with it.
		}
	}
	wg.Add(2) // we have 2 goroutine
	go sender()
	go receiver()
	b.StartTimer() // Here we begin the performance timer.

	close(begin) // Here we tell the two goroutines to begin.
	wg.Wait()    // wait for 2 goroutine to done
}
