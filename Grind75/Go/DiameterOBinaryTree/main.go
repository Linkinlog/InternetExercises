package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1+max(height(root.Left), height(root.Right))
}

func height(t *TreeNode) int {
	if t == nil {
		return 0
	}

	left := height(t.Left)
	right := height(t.Right)

	tallest := left
	if right > tallest {
		tallest = right
	}
	return 1 + tallest
}
