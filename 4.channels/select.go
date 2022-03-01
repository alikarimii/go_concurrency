package channels

import (
	"fmt"
	"time"
)

// You can see that it ran the default statement almost instantaneously.
// This allows you to exit a select block without blocking.
// Usually you’ll see a default clause used in conjunction with a for-select loop.
// This allows a goroutine to make progress on work while waiting for another
// goroutine to report a result
func RunSelect() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
