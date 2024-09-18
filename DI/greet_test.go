package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Omar")

	got := buffer.String()
	want := "Hello, Omar\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
