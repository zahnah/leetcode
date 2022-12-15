package leetcode

import "testing"

type FibTest struct {
	arg1     int
	expected int
}

var FibTests = []FibTest{
	{arg1: 2, expected: 1},
	{arg1: 3, expected: 2},
	{arg1: 4, expected: 3},
}

func TestFib(t *testing.T) {
	for _, test := range FibTests {
		result := fib(test.arg1)
		if result != test.expected {
			t.Errorf("Output %v not equal to expected %v", result, test.expected)
		}
	}
}

var ClimbStairsTests = []FibTest{
	{arg1: 2, expected: 2},
	{arg1: 3, expected: 3},
	{arg1: 4, expected: 5},
}

func TestClimbStairs(t *testing.T) {
	for _, test := range ClimbStairsTests {
		result := climbStairs(test.arg1)
		if result != test.expected {
			t.Errorf("Output %v not equal to expected %v", result, test.expected)
		}
	}
}
