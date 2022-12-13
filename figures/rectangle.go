package figures

type Rectangle struct {
	x float64
	y float64
}

func (rectangle Rectangle) Type() string {
	return "rectangle"
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.x * rectangle.y
}
