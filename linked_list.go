package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// MiddleNode
// 876. Middle of the Linked List
// https://leetcode.com/problems/middle-of-the-linked-list/
func MiddleNode(head *ListNode) *ListNode {
	var middle = head
	for head != nil && head.Next != nil {
		middle = middle.Next
		head = head.Next.Next
	}
	return middle
}
