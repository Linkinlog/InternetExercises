package main

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	tests := map[string]struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		"Case 1": {
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			p: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			q: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
			want: &TreeNode{
				Val: 6,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := lowestCommonAncestor(tc.root, tc.p, tc.q)
			if got == nil || got.Val != tc.want.Val {
				t.Errorf("got: %v, want: %v", got, tc.want.Val)
			}
		})
	}
}
