package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Scott")

	got := buffer.String()
	want := "Hello, Scott"

	if got != want {
		t.Errorf("got %q,  wanted %q", got, want)
	}
}
