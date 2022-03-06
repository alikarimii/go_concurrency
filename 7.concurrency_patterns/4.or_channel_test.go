package concurrencypatterns_test

import (
	"testing"
	"time"

	cp "github.com/alikarimii/go_concurrency/7.concurrency_patterns"
)

func TestOr(t *testing.T) {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-cp.OrChannel(
		sig(1*time.Second),
		sig(1*time.Minute),
		sig(5*time.Minute),
		sig(1*time.Hour),
		sig(2*time.Hour),
	)
	if time.Since(start).Round(1*time.Second) != 1*time.Second {
		t.Errorf("must done %v", time.Since(start))
	}
}
