package main

import (
	"testing"
)

func assertMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("Saying hello to someone", func(t *testing.T) {
		got := Hello("Lula")
		want := "Hello, Lula"

		assertMessage(t, got, want)
	})

	t.Run("Empty string defaults to 'Wold'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertMessage(t, got, want)
	})
}
