package leetcode

import (
	"fmt"
	"testing"
)

type MiddleNodeTest struct {
	arg1     *ListNode
	expected int
}

var MiddleNodeTests = []MiddleNodeTest{
	{arg1: createListNode([]int{0, 1, 2, 3, 4, 5}), expected: 3},
}

// ReverseString
func TestMiddleNode(t *testing.T) {
	for _, test := range MiddleNodeTests {
		result := MiddleNode(test.arg1)
		if result.Val != test.expected {
			t.Errorf("Output %q not equal to expected %q", result.Val, test.expected)
		}
	}
}

type RemoveNthFromEndTest struct {
	arg1     *ListNode
	arg2     int
	expected string
}

var RemoveNthFromEndTests = []RemoveNthFromEndTest{
	{arg1: createListNode([]int{0, 1, 2, 3, 4, 5}), arg2: 2, expected: "[0 1 2 3 5]"},
	{arg1: createListNode([]int{0, 1, 2, 3, 4, 5}), arg2: 3, expected: "[0 1 2 4 5]"},
	{arg1: createListNode([]int{1}), arg2: 1, expected: "[]"},
	{arg1: createListNode([]int{1, 2}), arg2: 1, expected: "[1]"},
	{arg1: createListNode([]int{1, 2}), arg2: 2, expected: "[2]"},
}

func TestRemoveNthFromEnd(t *testing.T) {
	for _, test := range RemoveNthFromEndTests {
		result := RemoveNthFromEnd(test.arg1, test.arg2)
		if fmt.Sprint(result) != test.expected {
			t.Errorf("Output %q not equal to expected %q", result, test.expected)
		}
	}
}

type IsValidBSTTest struct {
	arg1     *TreeNode
	expected bool
}

var IsValidBSTTests = []IsValidBSTTest{
	{arg1: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}, expected: true},
	{arg1: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}, expected: false},
	{arg1: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: nil,
	}, expected: false},
	{arg1: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}, expected: false},
	{arg1: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}, expected: false},
}

func TestIsValidBST(t *testing.T) {
	for i, test := range IsValidBSTTests {
		result := IsValidBST(test.arg1)
		if result != test.expected {
			t.Errorf("%d: Output %v not equal to expected %v", i, result, test.expected)
		}
	}
}

type LowestCommonAncestorTest struct {
	arg1     *TreeNode
	arg2     *TreeNode
	arg3     *TreeNode
	expected *TreeNode
}

var LowestCommonAncestorTest1 = &TreeNode{
	Val: 2,
	Left: &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	},
	Right: &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	},
}
var LowestCommonAncestorTest2 = &TreeNode{
	Val: 6,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	},
	Right: &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
	},
}
var LowestCommonAncestorTests = []LowestCommonAncestorTest{
	{arg1: LowestCommonAncestorTest1, arg2: LowestCommonAncestorTest1.Left, arg3: LowestCommonAncestorTest1.Right, expected: LowestCommonAncestorTest1},
	{arg1: LowestCommonAncestorTest2, arg2: LowestCommonAncestorTest2.Left, arg3: LowestCommonAncestorTest2.Right, expected: LowestCommonAncestorTest2},
	{arg1: LowestCommonAncestorTest2, arg2: LowestCommonAncestorTest2.Left, arg3: LowestCommonAncestorTest2.Left.Right, expected: LowestCommonAncestorTest2.Left},
	{arg1: LowestCommonAncestorTest2, arg2: LowestCommonAncestorTest2.Right, arg3: LowestCommonAncestorTest2.Right.Right, expected: LowestCommonAncestorTest2.Right},
	{arg1: LowestCommonAncestorTest2, arg2: LowestCommonAncestorTest2.Left, arg3: LowestCommonAncestorTest2.Left.Right, expected: LowestCommonAncestorTest2.Left},
	{arg1: LowestCommonAncestorTest2, arg2: LowestCommonAncestorTest2.Right.Left, arg3: LowestCommonAncestorTest2.Right.Right, expected: LowestCommonAncestorTest2.Right},
}

func TestLowestCommonAncestor(t *testing.T) {
	for i, test := range LowestCommonAncestorTests {
		result := LowestCommonAncestor(test.arg1, test.arg2, test.arg3)
		if result != test.expected {
			t.Errorf("%d: Output %v not equal to expected %v", i, result, test.expected)
		}
	}
}
