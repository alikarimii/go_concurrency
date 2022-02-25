package mutex_test

import (
	"sync"
	"testing"

	"github.com/alikarimii/go_concurrency/mutex"
)

func TestIncrement(t *testing.T) {
	myCounter := mutex.NewCounter()
	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			myCounter.Increment()
		}()
	}
	arithmetic.Wait()
	got := myCounter.GetCount()
	if got != 5 {
		t.Errorf("test failed, expect:%d , but got: %d", 5, got)
	}
}

func TestDecrement(t *testing.T) {
	myCounter := mutex.NewCounter()
	var arithmetic sync.WaitGroup
	// Decrement
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			myCounter.Decrement()
		}()
	}
	arithmetic.Wait()
	got := myCounter.GetCount()
	if got != -5 {
		t.Errorf("test failed, expect:%d , but got: %d", -5, got)
	}
}
