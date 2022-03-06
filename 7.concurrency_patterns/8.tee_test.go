package concurrencypatterns_test

import (
	"testing"

	cp "github.com/alikarimii/go_concurrency/7.concurrency_patterns"
)

func TestTee(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	out1, out2 := cp.Tee(done, cp.Take(done, cp.Repeat(done, 1, 2), 4))

	for val1 := range out1 {
		t.Logf("out1: %v, out2: %v\n", val1, <-out2)
	}
}
