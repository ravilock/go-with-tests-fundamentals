package Iteration

import (
	"testing"
)

func assertMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("Repeat a char 5 times", func (t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		
		assertMessage(t, repeated, expected)
	})
}