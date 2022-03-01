package channels_test

import (
	"testing"

	channels "github.com/alikarimii/go_concurrency/4.channels"
)

func TestReadNilChannel(t *testing.T) {
	readNil := func() {
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("[readNil]: %s", err)
			}
		}()
		channels.MustBlockedReadFromNilChannel()
	}
	readNil()
}

func TestWriteToNilChannel(t *testing.T) {
	func() {
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("[writeNil]: %s", err)
			}
		}()
		channels.MustBlockedWriteToNilChannel()
	}()
}

func TestCloseNilChannel(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Logf("[closeNil]: %s", err)
		} else {
			t.Error("this means close nil channel not panic and test failed")
		}
	}()
	channels.MustPanicCloseNilChannel()
}
