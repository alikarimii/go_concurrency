package channels_test

import (
	"testing"

	channels "github.com/alikarimii/go_concurrency/4.channels"
)

func TestSelect(t *testing.T) {
	channels.RunSelect()
}
