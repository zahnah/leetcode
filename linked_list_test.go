package leetcode

import (
	"fmt"
	"testing"
)

func ToString(node *ListNode) string {
	result := ""
	for node != nil {
		result = result + fmt.Sprint(node.Val)
		node = node.Next
	}
	return result
}

type MiddleNodeTest struct {
	arg1     *ListNode
	expected int
}

func getNode() *ListNode {
	return &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}
}

var MiddleNodeTests = []MiddleNodeTest{
	{arg1: getNode(), expected: 3},
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
	{arg1: getNode(), arg2: 2, expected: "01235"},
	{arg1: getNode(), arg2: 3, expected: "01245"},
	{arg1: &ListNode{
		Val: 1,
	}, arg2: 1, expected: ""},
	{arg1: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}, arg2: 1, expected: "1"},
	{arg1: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}, arg2: 2, expected: "2"},
}

func TestRemoveNthFromEnd(t *testing.T) {
	for _, test := range RemoveNthFromEndTests {
		result := RemoveNthFromEnd(test.arg1, test.arg2)
		if ToString(result) != test.expected {
			t.Errorf("Output %q not equal to expected %q", ToString(result), test.expected)
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

// [5,4,6,null,null,3,7]

func TestIsValidBST(t *testing.T) {
	for i, test := range IsValidBSTTests {
		result := IsValidBST(test.arg1)
		if result != test.expected {
			t.Errorf("%d: Output %v not equal to expected %v", i, result, test.expected)
		}
	}
}
