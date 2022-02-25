package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func Run() {
	var wg sync.WaitGroup
	// Here we call Add with an argument of 1 to
	// indicate that one goroutine is beginning
	wg.Add(1)
	go func() {
		// Here we call Done using the defer keyword to ensure that before
		// we exit the goroutine’s closure,
		// we indicate to the WaitGroup that we’ve exited
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1 * time.Second)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2 * time.Second)
	}()
	// Here we call Wait,
	// which will block the main goroutine
	// until all goroutines have indicated they have exited
	wg.Wait()
	fmt.Println("All goroutines complete.")
}
