package Sum

import (
	"testing"
)

func assertMessage(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("expected '%d' got'%d'", want, got)
	}
}

func TestSum(t *testing.T) {
	t.Run("Sum collection of any length", func(t *testing.T) {
		given := []int{1, 2, 3}

		got := Sum(given)
		want := 6

		assertMessage(t, got, want)
	})
}
