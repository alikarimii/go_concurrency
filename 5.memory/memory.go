package memory

import (
	"fmt"
	"runtime"
	"sync"
)

func Run() {
	var c <-chan interface{}
	var wg sync.WaitGroup
	// We require a goroutine that will never exit
	// so that we can keep a number of them in memory for measurement.
	noop := func() { wg.Done(); <-c }
	const numGoroutines = 1e4 // Here we define the number of goroutines to create.
	wg.Add(numGoroutines)
	// Here we measure the amount of memory consumed before creating our gorou‐ tines
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	// here we measure the amount of memory consumed after creating our goroutines
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}

func memConsumed() uint64 {
	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}
