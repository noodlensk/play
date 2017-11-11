package main

import (
	"fmt"
)

// Node is node of BST
type Node struct {
	left  *Node
	right *Node
	val   int
	sum   int
	dup   int
}

func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	var root *Node
	for i := len(nums) - 1; i >= 0; i-- {
		root = insert(nums[i], root, res, i, 0)
	}
	return res
}
func insert(num int, node *Node, ans []int, i, preSum int) *Node {
	if node == nil {
		node = &Node{
			val: num,
			sum: 0,
			dup: 1,
		}
		ans[i] = preSum
	} else if node.val == num {
		node.dup++
		ans[i] = preSum + node.sum
	} else if node.val > num {
		node.sum++
		node.left = insert(num, node.left, ans, i, preSum)
	} else {
		node.right = insert(num, node.right, ans, i, preSum+node.dup+node.sum)
	}
	return node
}
func main() {
	res := countSmaller([]int{5, 2, 6, 1})
	fmt.Printf("%+v\n", res)
}
