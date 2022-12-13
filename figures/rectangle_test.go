package figures

import (
	"testing"
)

var rectangle Figure = Rectangle{
	x: 4,
	y: 3,
}

func TestRectangle_Type(t *testing.T) {
	if rectangle.Type() != "rectangle" {
		t.Errorf("Output %v not equal to expected %v", rectangle.Type(), "rectangle")
	}
}

func TestRectangle_Area(t *testing.T) {
	if rectangle.Area() != float64(12) {
		t.Errorf("Output %v not equal to expected %v", rectangle.Area(), 12)
	}
}
