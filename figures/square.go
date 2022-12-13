package figures

type Square struct {
	x float64
}

func (square Square) Type() string {
	return "square"
}

func (square Square) Area() float64 {
	return square.x * square.x
}
