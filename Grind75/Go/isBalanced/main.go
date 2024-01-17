package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	// recurse left/right nodes, if balanced, increment counter
	// if lcounter - rcounter between 0-2
	if root == nil {
		return true
	}
	var loop func(r *TreeNode) int
	loop = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		lh := loop(r.Left)
		if lh == -1 {
			return -1
		}
		rh := loop(r.Right)
		if rh == -1 {
			return -1
		}

		if math.Abs(float64(lh)-float64(rh)) > 1 {
			return -1
		} else {
			return int(math.Max(float64(lh), float64(rh)) + 1)
		}
	}
	val := loop(root)
	return val > 0
}

