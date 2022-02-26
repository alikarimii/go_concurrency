package mutex_test

import (
	"sync"
	"testing"

	mutex "github.com/alikarimii/go_concurrency/3.mutex"
)

func TestOnce(t *testing.T) {
	c := mutex.NewCount()
	var increments sync.WaitGroup
	var once sync.Once

	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(c.Increment)
		}()
	}
	if c.Value() != 1 {
		t.Error("value must be 1")
	}
	t.Log("current value is: ", c.Value())
}
