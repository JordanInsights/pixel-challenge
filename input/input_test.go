// package input_test tests input package functions
package input_test

import (
	"bytes"
	"io"
	"io/ioutil"
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

func TestGetFilepath(t *testing.T) {
	var got, want string
	want = "./test-filepath.raw"

	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, want)
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatal(err)
	}

	got = input.GetFilepath(in)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
