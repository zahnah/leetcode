package leetcode

import (
	"testing"
)

type SearchTest struct {
	arg1     []int
	arg2     int
	expected int
}

var SearchTests = []SearchTest{
	{arg1: []int{-1, 0, 3, 5, 9, 12}, arg2: 9, expected: 4},
	{arg1: []int{-1, 0, 3, 5, 9, 12}, arg2: 2, expected: -1},
}

// ReverseString
func TestSearch(t *testing.T) {
	for _, test := range SearchTests {
		if result := Search(test.arg1, test.arg2); result != test.expected {
			t.Errorf("Output %q not equal to expected %q", result, test.expected)
		}
	}
}

// ReverseString
func TestFirstBadVersion(t *testing.T) {
	if result := FirstBadVersion(5); result != 4 {
		t.Errorf("Output %q not equal to expected %q", result, 4)
	}
}
