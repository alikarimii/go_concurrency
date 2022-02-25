package mutex

import (
	"fmt"
	"sync"
)

func NewCounter() counter {
	return counter{lock: sync.Mutex{}, count: 0}
}

type counter struct {
	lock  sync.Mutex
	count int
}

func (c *counter) Increment() {
	// in this case the count variable guarded by a Mutex, lock
	c.lock.Lock()
	// Here we indicate that we’re done with the critical section lock is guarding.
	defer c.lock.Unlock()
	c.count++
	fmt.Printf("Incrementing: %d\n", c.count)
}

func (c *counter) Decrement() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.count--
	fmt.Printf("Decrementing: %d\n", c.count)
}

// for testing
func (c *counter) GetCount() int {
	return c.count
}
