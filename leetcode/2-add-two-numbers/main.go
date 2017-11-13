package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	sum(l1, l2, root)
	return root
}
func sum(l1 *ListNode, l2 *ListNode, res *ListNode) *ListNode {
	var l1Next, l2Next *ListNode
	if l1 != nil {
		res.Val = res.Val + l1.Val
		l1Next = l1.Next
	}
	if l2 != nil {
		res.Val = res.Val + l2.Val
		l2Next = l2.Next
	}

	if res.Val > 9 {
		res.Val = res.Val - 10
		res.Next = &ListNode{Val: 1}
	}

	if (l1Next != nil) || (l2Next != nil) {
		if res.Next == nil {
			res.Next = &ListNode{}
		}
		return sum(l1Next, l2Next, res.Next)
	}

	return res
}

func main() {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 6,
			},
		},
	}
	l2 := &ListNode{
		Val: 0,
	}
	fmt.Printf("%+v\n", addTwoNumbers(l1, l2))
}
