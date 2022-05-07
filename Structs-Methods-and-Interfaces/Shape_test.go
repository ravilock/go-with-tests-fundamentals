package Shapes

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{4.0, 10.0}, 28.0},
		{Circle{50.0}, 314.1592653589793},
		{Triangle{3.0, 3.0, 3.0}, 9.0},
	}
	for _, testTable := range perimeterTests {
		checkPerimeter(t, testTable.shape, testTable.want)
	}
}

func checkPerimeter(t testing.TB, shape Shape, want float64) {
	t.Helper()

	got := shape.Perimeter()
	if got != want {
		t.Errorf("%#v got %g want %g", shape, got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{4.0, 10.0}, want: 40.0},
		{shape: Circle{10.0}, want: 314.1592653589793},
		{shape: Triangle{3.0, 4.0, 5.0}, want: 6.0},
	}
	for _, testTable := range areaTests {
		checkArea(t, testTable.shape, testTable.want)
	}
}

func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()

	got := shape.Area()
	if got != want {
		t.Errorf("%#v got %g want %g", shape, got, want)
	}
}
