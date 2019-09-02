package shapes

import (
	"math"
)

type Rectangle struct {
	Length  float64
	Breadth float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
