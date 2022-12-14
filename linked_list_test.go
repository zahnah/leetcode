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
	{arg1: createTreeNode([]int{2, 1, 3}), expected: true},
	{arg1: createTreeNode([]int{5, 1, 4, -1, -1, 3, 6}), expected: false},
	{arg1: createTreeNode([]int{5, 1, -1, 3, -1}), expected: false},
	{arg1: createTreeNode([]int{2, 2, 2}), expected: false},
	{arg1: createTreeNode([]int{5, 4, 6, -1, -1, 3, 7}), expected: false},
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

var LowestCommonAncestorTest1 = createTreeNode([]int{2, 1, 3})
var LowestCommonAncestorTest2 = createTreeNode([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5})
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

type TreeNodeStringTest struct {
	arg1     *TreeNode
	expected string
}

var TreeNodeStringTests = []TreeNodeStringTest{
	{arg1: createTreeNode([]int{2, 1, 3}), expected: "[2 1 3]"},
	{arg1: createTreeNode([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}), expected: "[6 2 8 0 4 7 9 nil nil 3 5]"},
	{arg1: createTreeNode([]int{6, 2, 8, 1, 4, -1, 10, 0, -1, 3, 5, -1, -1, -1, 12}), expected: "[6 2 8 1 4 nil 10 0 nil 3 5 nil nil nil 12]"},
}

func TestTreeNodeString(t *testing.T) {
	for i, test := range TreeNodeStringTests {
		if fmt.Sprint(test.arg1) != test.expected {
			t.Errorf("%d: Output %v not equal to expected %v", i, test.arg1, test.expected)
		}
	}
}

type CreateListNodeTest struct {
	arg1     []int
	expected string
}

var CreateListNodeTests = []CreateListNodeTest{
	{arg1: []int{2, 1, 3}, expected: "[2 1 3]"},
	{arg1: []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}, expected: "[6 2 8 0 4 7 9 nil nil 3 5]"},
	{arg1: []int{6, -1, 8, -1, -1, 7, 9, -1, -1, -1, -1, -1, -1, -1, 12}, expected: "[6 nil 8 nil nil 7 9 nil nil nil nil nil nil nil 12]"},
}

func TestCreateListNode(t *testing.T) {
	for i, test := range CreateListNodeTests {
		if fmt.Sprint(createTreeNode(test.arg1)) != test.expected {
			t.Errorf("%d: Output %v not equal to expected %v", i, createTreeNode(test.arg1), test.expected)
		}
	}
}
