package leetcode

import (
	"reflect"
	"testing"
)

type FloodFillTest struct {
	image    [][]int
	sr       int
	sc       int
	color    int
	expected [][]int
}

var FloodFillTests = []FloodFillTest{
	{
		image:    [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
		sr:       1,
		sc:       1,
		color:    2,
		expected: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
	},
	{
		image:    [][]int{{0, 0, 0}, {0, 0, 0}},
		sr:       0,
		sc:       0,
		color:    0,
		expected: [][]int{{0, 0, 0}, {0, 0, 0}},
	},
}

// ReverseString
func TestFloodFill(t *testing.T) {
	for _, test := range FloodFillTests {
		result := floodFill(test.image, test.sr, test.sc, test.color)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Output %v not equal to expected %v", result, test.expected)
		}
	}
}
