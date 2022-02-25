package mutex

import (
	"fmt"
	"sync"
	"time"
)

// we have a queue of fixed length 2, and 10 items
// we want to push onto the queue.
// We want to enqueue items as soon as there is room,
// so we want to be notified as soon as there’s room in the queue
func NewQueue() myQueue {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)
	return myQueue{c, queue}
}

type myQueue struct {
	c     *sync.Cond
	queue []interface{}
}

func (q *myQueue) RemoveFromQueue(delay time.Duration) {
	time.Sleep(delay)
	q.c.L.Lock()
	q.queue = q.queue[1:]
	fmt.Println("Removed from queue")
	q.c.L.Unlock()
	// Here we let a goroutine waiting on the condition know that something has occurred
	q.c.Signal()
}

func (q *myQueue) AddToQueue(job interface{}) {
	// We enter the critical section for the condition
	// by calling Lock on the condition’sLocker
	q.c.L.Lock()
	// Here we check the length of the queue in a loop.
	// This is important because a signal on
	// the condition doesn’t necessarily mean what you’ve been waiting
	// for has occurred—only that something has occurred
	for len(q.queue) == 2 {
		// We call Wait, which will suspend the main goroutine
		// until a signal on the condition has been sent
		q.c.Wait()
	}
	fmt.Println("Adding to queue")
	q.queue = append(q.queue, job)
	q.c.L.Unlock()
}

func (q *myQueue) QueueLen() int {
	return len(q.queue)
}
