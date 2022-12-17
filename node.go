package leetcode

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (node *Node) String() string {
	var result []string
	q1 := []*Node{node}
	var q2 []*Node
	for len(q1) > 0 {
		for _, child := range q1 {
			val := "nil"
			if child != nil {
				val = fmt.Sprint(child.Val)
				if child.Next != nil {
					val += "->" + fmt.Sprint(child.Next.Val)
				}
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
			q2 = []*Node{}
		} else {
			break
		}
	}
	return fmt.Sprint(result)
}

func createNode(values []int) *Node {
	i := 1
	deep := 2
	root := Node{
		Val: values[0],
	}
	q1 := []*Node{&root}
	var q2 []*Node
	for i < len(values) {
		for j := 0; j < deep && i < len(values); j++ {
			ji := j / 2
			if values[i] != -1 {
				node := &Node{
					Val: values[i],
				}
				if j%2 == 0 {
					q1[ji].Left = node
				} else {
					q1[ji].Right = node
				}
			}
			i++
		}
		for _, child := range q1 {
			if child != nil {
				q2 = append(q2, child.Left, child.Right)
			} else {
				q2 = append(q2, nil, nil)
			}
		}
		q1 = q2
		q2 = []*Node{}
		deep *= 2
	}
	return &root
}

// connect
// 116. Populating Next Right Pointers in Each Node
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
func connect(root *Node) *Node {
	nodes := []*Node{root}
	var nextNodes []*Node
	for len(nodes) > 0 {
		for i := 0; i < len(nodes); i++ {
			if i < len(nodes)-1 {
				nodes[i].Next = nodes[i+1]
			}

			if nodes[i].Left != nil {
				nextNodes = append(nextNodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nextNodes = append(nextNodes, nodes[i].Right)
			}

		}

		nodes = nextNodes
		nextNodes = []*Node{}
	}
	return root
}
