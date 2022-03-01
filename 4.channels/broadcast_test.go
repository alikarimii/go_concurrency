package channels_test

import (
	"sync"
	"testing"

	channels "github.com/alikarimii/go_concurrency/4.channels"
)

func TestBroadcast(t *testing.T) {

	btn := channels.NewButton()
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
