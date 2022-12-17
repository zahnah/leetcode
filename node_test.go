package leetcode

import (
	"fmt"
	"testing"
)

type ConnectTest struct {
	arg1     *Node
	expected string
}

var ConnectTests = []ConnectTest{
	{arg1: createNode([]int{1, 2, 3, 4, 5, 6, 7}), expected: "[1 2->3 3 4->5 5->6 6->7 7]"},
}

// ReverseString
func TestConnect(t *testing.T) {
	for _, test := range ConnectTests {
		result := connect(test.arg1)
		if fmt.Sprint(result) != test.expected {
			t.Errorf("Output %q not equal to expected %q", result, test.expected)
		}
	}
}
