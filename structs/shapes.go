package structs

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Length + rectangle.Breadth)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Length * rectangle.Breadth
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}

func (triangle Triangle) Area() float64 {
	return 0.5 * triangle.Height * triangle.Base
}
