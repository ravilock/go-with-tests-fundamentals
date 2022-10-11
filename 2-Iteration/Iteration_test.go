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
	t.Run("Repeat a char N times", func(t *testing.T) {
		n := 15
		repeated := Repeat("a", n)
		expected := "aaaaaaaaaaaaaaa"

		assertMessage(t, repeated, expected)
	})

	t.Run("If the repeatCount is less than 1, it should return an empty string", func(t *testing.T) {
		n := 0
		repeated := Repeat("a", n)
		expected := ""

		assertMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Repeat("a", 15)
	}
}
