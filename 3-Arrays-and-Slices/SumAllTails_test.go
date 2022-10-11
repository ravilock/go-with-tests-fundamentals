package Sum

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	t.Run("Sum the tails of a slice of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected '%d' got '%d'", want, got)
		}
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected '%d' got '%d'", want, got)
		}
	})

}

func BenchmarkSumAllTails(t *testing.B) {
	for n := 0; n < t.N; n++ {
		SumAllTails([]int{1, 2, 3, 4}, []int{6, 9, 0, 4, 3, 1})
	}
}
