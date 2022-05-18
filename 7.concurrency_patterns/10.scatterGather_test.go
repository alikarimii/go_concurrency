package concurrencypatterns_test

import (
	"testing"

	concurrencypatterns "github.com/alikarimii/go_concurrency/7.concurrency_patterns"
)

func TestScatterGather(t *testing.T) {
	jobs := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	routines := 3
	// this can return some chan and error to get result
	concurrencypatterns.ScatterGather(routines, jobs)
}
