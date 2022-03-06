package concurrencypatterns_test

import (
	"testing"

	cp "github.com/alikarimii/go_concurrency/7.concurrency_patterns"
)

func TestPipeline(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	total := 0
	intStream := cp.Generator(done, 1, 2, 3)
	pipeline := cp.Add(done, cp.Multiply(done, intStream, 2), 1)
	for v := range pipeline {
		total += v
	}
	if total != 15 {
		t.Errorf("total must be %d but is: %d", 15, total)
	}
}
