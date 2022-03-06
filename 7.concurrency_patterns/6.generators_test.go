package concurrencypatterns_test

import (
	"math/rand"
	"testing"

	cp "github.com/alikarimii/go_concurrency/7.concurrency_patterns"
)

func TestTakeAndRepeat(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	total := 0
	for num := range cp.Take(done, cp.Repeat(done, 1), 10) {
		total += num.(int)
	}
	if total != 10 {
		t.Errorf("total must be %d but is: %d", 10, total)
	}
}

func TestRepeatFn(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	rand := func() interface{} { return rand.Int() }
	for num := range cp.Take(done, cp.RepeatFn(done, rand), 10) {
		t.Log(num)
	}
}
