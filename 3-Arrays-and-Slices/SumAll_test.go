package Sum

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("Sum a slice of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected '%d' got'%d'", want, got)
		}
	})
}

func BenchmarkSumAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAll([]int{1, 2, 3, 4}, []int{6, 9, 0, 4, 3, 1})
	}
}
