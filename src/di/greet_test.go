package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Steve")

	got := buffer.String()
	want := "Hello, Steve"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
