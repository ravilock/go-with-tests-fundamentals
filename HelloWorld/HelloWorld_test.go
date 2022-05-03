package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Lula")
	want := "Hello, Lula"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
