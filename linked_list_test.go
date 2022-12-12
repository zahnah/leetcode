package leetcode

import "testing"

type MiddleNodeTest struct {
	arg1     *ListNode
	expected int
}

var MiddleNodeTests = []MiddleNodeTest{
	{arg1: &ListNode{
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
	}, expected: 3},
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
