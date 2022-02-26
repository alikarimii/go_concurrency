package mutex_test

import (
	"testing"

	mutex "github.com/alikarimii/go_concurrency/3.mutex"
)

func TestCond(t *testing.T) {

	my := mutex.NewQueue()
	if my.QueueLen() != 0 {
		t.Fatal("initial queue len must be zero")
	}

	my.AddToQueue("every thing")

	if my.QueueLen() != 1 {
		t.Fatal("queue len must be 1")
	}

	my.RemoveFromQueue(0)
	if my.QueueLen() != 0 {
		t.Fatal("queue len must be zero")
	}
	// AddToQueue will be blocked if we are added more than 2 item
	// to queue without remove from them
	for i := 0; i < 10; i++ {
		my.AddToQueue(i)
		go my.RemoveFromQueue(1)
	}

}
