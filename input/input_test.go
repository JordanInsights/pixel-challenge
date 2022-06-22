// package input_test tests input package functions
package input_test

import (
	"bytes"
	"pixel-challenge/input"
	"testing"
)

func TestPrintMessage(t *testing.T) {
	buffer := bytes.Buffer{}
	input.PrintMessage(&buffer, "Hello, and welcome to pixel comparison!")

	got := buffer.String()
	want := "Hello, and welcome to pixel comparison!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
