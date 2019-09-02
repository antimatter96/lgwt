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

func Perimeter(rectangle Rectangle) (perimeter float64) {
	perimeter = 2 * (rectangle.Length + rectangle.Breadth)
	return
}

func Area(rectangle Rectangle) (perimeter float64) {
	perimeter = rectangle.Length * rectangle.Breadth
	return
}

func Area(circle Circle) (perimeter float64) {
	perimeter = circle.Radius * circle.Radius * math.Pi
	return
}
