package channels

import (
	"sync"
)

func NewButton() button {
	return button{c: make(chan struct{})}
}

type button struct {
	c chan struct{}
}

func (b *button) Broadcast() {
	close(b.c)
}
func (b *button) Subscribe(fn func()) {
	var goroutineRunning sync.WaitGroup
	goroutineRunning.Add(1)
	go func() {
		goroutineRunning.Done()
		<-b.c
		fn()
	}()
	goroutineRunning.Wait()
}
