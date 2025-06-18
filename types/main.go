package main

import "fmt"

const (
	ONE   = 1
	TWO   = 2
	THREE = 3
)

type TreeNode struct {
	left, right *TreeNode
	value       any
}
type ListNode struct {
	value any
	next  *ListNode
}

func main() {
	println(ONE)

	root := TreeNode{
		left:  nil,
		right: nil,
		value: 1,
	}
	leftNode := TreeNode{
		left : nil,
		right: nil,
		value: 2,
	}
	root.left = &leftNode
	root.right = &TreeNode{
		left:  nil,
		right: nil,
		value: 3,
	}
	fmt.Println(root)
}
