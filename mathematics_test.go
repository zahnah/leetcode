package leetcode

import (
	"fmt"
	"testing"
)

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

type CombineTest struct {
	arg1     int
	arg2     int
	expected [][]int
}

var CombineTests = []CombineTest{
	{arg1: 1, arg2: 1, expected: [][]int{{1}}},
	{arg1: 4, arg2: 2, expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	{arg1: 5, arg2: 4, expected: [][]int{{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5}}},
}

func TestCombine(t *testing.T) {
	for _, test := range CombineTests {
		result := combine(test.arg1, test.arg2)
		if fmt.Sprint(result) != fmt.Sprint(test.expected) {
			t.Errorf("Output %v not equal to expected %v", result, test.expected)
		}
	}
}
