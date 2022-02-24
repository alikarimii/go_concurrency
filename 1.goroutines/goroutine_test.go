package goroutines_test

import (
	"testing"

	"github.com/alikarimii/go_concurrency/goroutines"
)

func TestRun(t *testing.T) {
	goroutines.Run()
}

func BenchmarkRun(b *testing.B) {
	goroutines.Run()
}
