package Shapes

import (
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Side1, Side2, Side3 float64
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

// Using Heron's Formula
func (triangle Triangle) Area() float64 {
	factor1 := triangle.Side1 + triangle.Side2 + triangle.Side3
	factor2 := -triangle.Side1 + triangle.Side2 + triangle.Side3
	factor3 := triangle.Side1 - triangle.Side2 + triangle.Side3
	factor4 := triangle.Side1 + triangle.Side2 - triangle.Side3
	product := factor1 * factor2 * factor3 * factor4
	return math.Sqrt(product) / 4
}

func (triangle Triangle) Perimeter() float64 {
	return triangle.Side1 + triangle.Side2 + triangle.Side3
}
