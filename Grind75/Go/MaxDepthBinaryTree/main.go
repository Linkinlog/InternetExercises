package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lSum := maxDepth(root.Left)
	rSum := maxDepth(root.Right)

	tallest := lSum
	if rSum > tallest {
		tallest = rSum
	}

	return 1 + tallest
}
