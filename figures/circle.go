package figures

import "math"

type Circle struct {
	r float64
}

func (circle Circle) Area() float64 {
	return float64(2) * circle.r * math.Pi
}

func (circle Circle) Type() string {
	return "circle"
}
