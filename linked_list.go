package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	var result []rune
	each := node

	for each != nil {
		result = append(result, rune(each.Val))
		each = each.Next
	}
	return fmt.Sprint(result)
}

func createListNode(list []int) *ListNode {
	var root = ListNode{
		Val: list[0],
	}
	if len(list) > 1 {
		root.Next = createListNode(list[1:])
	}
	return &root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) String() string {
	var result []string
	q1 := []*TreeNode{node}
	var q2 []*TreeNode
	for len(q1) > 0 {
		for _, child := range q1 {
			val := "nil"
			if child != nil {
				val = fmt.Sprint(child.Val)
				q2 = append(q2, child.Left, child.Right)
			} else {
				q2 = append(q2, nil, nil)
			}
			result = append(result, val)
		}
		lastValueIndex := len(q2) - 1
		for ; lastValueIndex > 0; lastValueIndex -= 2 {
			if q2[lastValueIndex] != nil || q2[lastValueIndex-1] != nil {
				break
			}
		}
		if lastValueIndex > 0 {
			q2 = q2[0 : lastValueIndex+1]
			q1 = q2
			q2 = []*TreeNode{}
		} else {
			break
		}
	}
	return fmt.Sprint(result)
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

// RemoveNthFromEnd
// 19. Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	next := head
	end := head
	for i := 0; i < n && end != nil; i++ {
		end = end.Next
	}
	if end == nil {
		head = head.Next
		return head
	}
	for end.Next != nil {
		next = next.Next
		end = end.Next
	}
	next.Next = next.Next.Next
	return head
}

// IsValidBST
// 98. Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/
func IsValidBST(root *TreeNode) bool {
	return IsValidBSTChild(root, nil, nil)
}
func IsValidBSTChild(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if nil != min && root.Val >= min.Val || nil != max && root.Val <= max.Val {
		return false
	}
	if root.Left != nil && !IsValidBSTChild(root.Left, root, max) {
		return false
	}
	if root.Right != nil && !IsValidBSTChild(root.Right, min, root) {
		return false
	}
	return true
}

// LowestCommonAncestor
// 235. Lowest Common Ancestor of a Binary Search Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || (root.Val > p.Val || root.Val > q.Val) && (root.Val < q.Val || root.Val < p.Val) {
		return root
	} else if root.Left != nil && root.Val > p.Val && root.Val > q.Val {
		return LowestCommonAncestor(root.Left, p, q)
	} else {
		return LowestCommonAncestor(root.Right, p, q)
	}
}
