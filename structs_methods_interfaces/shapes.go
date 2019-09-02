package shapes

type Rectangle struct {
	length  float64
	breadth float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	perimeter = 2 * (rectangle.length + rectangle.breadth)
	return
}

func Area(rectangle Rectangle) (perimeter float64) {
	perimeter = rectangle.length * rectangle.breadth
	return
}
