package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	shape := Rectangle{10.0, 10.0}
	var want float64 = 40
	got := Perimeter(shape)

	if got != want {
		t.Errorf("got %.2f; hasArea %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// new option: table-driven tests
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Length: 5.5, Breadth: 2}, hasArea: 11.0},
		{name: "Circle", shape: Circle{Radius: 5}, hasArea: math.Pi * 5 * 5},
		{name: "Triangle", shape: Triangle{Height: 6, Base: 3}, hasArea: 9},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g, hasArea %g", tt.shape, got, tt.hasArea)
			}
		})
	}

	// old options
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f, hasArea %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{5.5, 2}
		checkArea(t, rectangle, 11.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{5}
		checkArea(t, circle, math.Pi*5*5)
	})
}
