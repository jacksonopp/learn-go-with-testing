package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("it should calculate the perimeter", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := Perimeter(rect)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	// checkArea := func(shape Shape, want float64, t testing.TB) {
	// 	t.Helper()

	// 	got := shape.Area()

	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("it should calculate the area of a rectangle", func(t *testing.T) {
	// 	rect := Rectangle{10.0, 10.0}
	// 	checkArea(rect, 100.0, t)
	// })

	// t.Run("it should calculate the area of a rectangle", func(t *testing.T) {
	// 	circ := Circle{10.0}
	// 	checkArea(circ, math.Pi*100, t)
	// })

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, math.Pi * 100},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}

}
