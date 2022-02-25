package mutex

import (
	"sync"
)

func NewButton() button {
	return button{c: sync.NewCond(&sync.Mutex{})}
}

type button struct {
	c *sync.Cond
}

func (b *button) Broadcast() {
	b.c.Broadcast()
}
func (b *button) Subscribe(fn func()) {
	var goroutineRunning sync.WaitGroup
	goroutineRunning.Add(1)
	go func() {
		goroutineRunning.Done()
		b.c.L.Lock()
		defer b.c.L.Unlock()
		b.c.Wait()
		fn()
	}()
	goroutineRunning.Wait()
}
