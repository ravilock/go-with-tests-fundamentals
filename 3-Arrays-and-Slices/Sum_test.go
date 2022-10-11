package Sum

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum collection of any length", func(t *testing.T) {
		given := []int{1, 2, 3}

		got := Sum(given)
		want := 6

		if got != want {
			t.Errorf("expected '%d' got'%d'", want, got)
		}
	})
}
