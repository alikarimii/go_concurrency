package concurrencypatterns_test

import (
	"testing"

	cp "github.com/alikarimii/go_concurrency/7.concurrency_patterns"
)

func TestFanInOut(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	in := cp.Generator(done, 1, 2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	// Fan-out part
	c1 := cp.Sq(done, in)
	c2 := cp.Sq(done, in)

	// Consume the merged output from c1 and c2.
	// // Fan-in part
	for n := range cp.FanInInt(done, c1, c2) {
		t.Logf("sq is: %d", n) // 1,4,9 but maybe not in this order
	}
}
