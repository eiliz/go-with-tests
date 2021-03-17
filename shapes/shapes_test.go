package shapes

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {

	r := Rectangle{10.0, 10.0}
	got := r.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("expected %g, got %g", want, got)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 6, Height: 12}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: math.Pi * math.Pow(10, 2)},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {

		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v expected %g, got %g", tt.shape, tt.hasArea, got)
			}
		})
	}
}

// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("expected %g, got %g", want, got)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		r := Rectangle{6, 12}
// 		checkArea(t, r, 72.0)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		c := Circle{10}
// 		checkArea(t, c, math.Pi*math.Pow(10, 2))
// 	})
// }
