package goroutines_test

import (
	"testing"

	goroutines "github.com/alikarimii/go_concurrency/1.goroutines"
)

func TestRun(t *testing.T) {
	goroutines.Run()
}

func BenchmarkRun(b *testing.B) {
	goroutines.Run()
}
