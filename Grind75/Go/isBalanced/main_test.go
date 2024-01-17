package main

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected bool
	}{
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{Val: 2},
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{Val: 3},
					Left: &TreeNode{
						Val: 3,
						Right: &TreeNode{Val: 4},
						Left: &TreeNode{
							Val:   4,
						},
					},
				},
			},
			expected: false,
		},
		{
			root: &TreeNode{
				Val: 69,
				Left: &TreeNode{
					Val: 420,
				},
				Right: &TreeNode{
					Val: 420,
					Right: &TreeNode{
						Val: 666,
						Right: &TreeNode{
							Val:   69,
							Right: &TreeNode{},
						},
					},
				},
			},
			expected: false,
		},
		{
			root: &TreeNode{
				Val: 69,
				Left: &TreeNode{
					Val: 420,
					Left: &TreeNode{
						Val:  666,
						Left: nil,
					},
				},
				Right: &TreeNode{
					Val: 420,
					Right: &TreeNode{
						Val: 666,
						Right: &TreeNode{
							Val: 69,
							Right: &TreeNode{
								Val: 420,
								Right: &TreeNode{
									Val: 666,
									Right: &TreeNode{
										Val: 69,
									},
								},
							},
						},
					},
				},
			},
			expected: false,
		},
		{
			root: &TreeNode{
				Val: 69,
				Left: &TreeNode{
					Val: 420,
					Left: &TreeNode{
						Val:  666,
						Left: nil,
					},
				},
				Right: &TreeNode{
					Val: 420,
					Right: &TreeNode{
						Val:   666,
						Right: nil,
					},
				},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		actual := isBalanced(test.root)
		if actual != test.expected {
			t.Errorf("isBalanced(%v) = %v; expected %v", test.root, actual, test.expected)
		}
	}
}
