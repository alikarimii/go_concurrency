package channels_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestSample(t *testing.T) {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)
	intStream := make(chan int, 3)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			intStream <- i
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
		}
	}()
	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}
