package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	name := "Carl"
	buffer := bytes.Buffer{}

	Greet(&buffer, name)

	got := buffer.String()
	want := "Hello, Carl!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
