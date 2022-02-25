package mutex_test

import (
	"sync"
	"testing"

	"github.com/alikarimii/go_concurrency/mutex"
)

func TestBroadcast(t *testing.T) {

	btn := mutex.NewButton()

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	btn.Subscribe(func() {
		t.Log("Maximizing window.")
		clickRegistered.Done()
	})
	btn.Subscribe(func() {
		t.Log("Displaying annoying dialog box!")
		clickRegistered.Done()
	})
	btn.Subscribe(func() {
		t.Log("Mouse clicked.")
		clickRegistered.Done()
	})
	btn.Broadcast()
	clickRegistered.Wait()
}
