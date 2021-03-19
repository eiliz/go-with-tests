package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Gra Gra")

	got := buffer.String()
	want := "Gra Gra"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
