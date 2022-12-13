package figures

import (
	"testing"
)

var square Figure = Square{
	x: 4,
}

func TestSquare_Type(t *testing.T) {
	if square.Type() != "square" {
		t.Errorf("Output %v not equal to expected %v", square.Type(), "square")
	}
}

func TestSquare_Area(t *testing.T) {
	if square.Area() != float64(16) {
		t.Errorf("Output %v not equal to expected %v", square.Area(), 16)
	}
}
