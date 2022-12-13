package figures

import (
	"testing"
)

var circle Figure = Circle{
	r: 4,
}

func TestCircle_Type(t *testing.T) {
	if circle.Type() != "circle" {
		t.Errorf("Output %v not equal to expected %v", circle.Type(), "rectangle")
	}
}

func TestCircle_Area(t *testing.T) {
	if circle.Area() != 25.132741228718345 {
		t.Errorf("Output %v not equal to expected %v", circle.Area(), 25.132741228718345)
	}
}
