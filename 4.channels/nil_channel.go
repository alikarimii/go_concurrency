package channels

import (
	"fmt"
	"log"
	"time"
)

func MustBlockedReadFromNilChannel() {
	var dataStream chan interface{}
	select {
	case <-dataStream: // dead lock happend
	case <-time.After(1 * time.Second):
		fmt.Println("every things ok, no dead lock")
		return
	}
	log.Fatal("if code reach here,means dataStream not locked and test failed")
}

func MustBlockedWriteToNilChannel() {
	var dataStream chan interface{}
	go func() {
		dataStream <- struct{}{} // dead lock happend
	}()

	select {
	case <-dataStream:
	case <-time.After(1 * time.Second):
		fmt.Println("every things ok, no dead lock")
		return
	}
	log.Fatal("if code reach here,means write to dataStream not locked and test failed")
}
func MustPanicCloseNilChannel() {
	var dataStream chan interface{}
	close(dataStream)
}
