package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{10.0, 10.0}, want: 40.0},
		{name: "circle", shape: Circle{10.0}, want: 62.831853071795862},
		{name: "triangle", shape: Triangle{3.0, 5.0, 6.0}, want: 14.0},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
			}
		})

	}

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{10.0, 9.0}, want: 90.0},
		{name: "circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{3.0, 5.0, 6.0}, want: 7.483314773547883},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
			}
		})
	}

}
