package problems_39254_test

import (
	"testing"

	dg "github.com/alikarimii/go_concurrency/example/dg/problems_39254"
)

func TestQ1(t *testing.T) {
	trueInput := []byte("3\n1 2\n3 4\n5 6\n2\n5")
	falseInput := []byte("2\n1 10\n10 20\n21\n21")
	parser := dg.NewParser()

	if !dg.Run(parser, trueInput) {
		t.Error("must be true")
	}
	if dg.Run(parser, falseInput) {
		t.Error("must be false")
	}
}
