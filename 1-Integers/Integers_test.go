package Integers

import (
	"testing"
)

func assertMessage(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("expected '%d' got'%d'", want, got)
	}
}

func TestAdd(t* testing.T) {
	t.Run("Sum 2 + 2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertMessage(t, expected, sum)
	})
}
