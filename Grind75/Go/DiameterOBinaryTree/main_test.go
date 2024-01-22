package main

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want int
	}{
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, 2},
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}, 3},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 6,
								Right: &TreeNode{
									Val: 8,
								},
							},
							Right: &TreeNode{
								Val: 7,
							},
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: 6,
		},
	}

	for _, c := range cases {
		got := diameterOfBinaryTree(c.root)
		if got != c.want {
			t.Errorf("diameterOfBinaryTree(%v) == %d, want %d", c.root, got, c.want)
		}
	}
}

func TestHeight(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want int
	}{
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, 2},
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}, 3},
		{&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 6,
								Left: &TreeNode{
									Val: 7,
									Right: &TreeNode{
										Val: 8,
										Right: &TreeNode{
											Val: 9,
										},
									},
								},
							},
						},
					},
				},
			},
		}, 9},
	}

	for _, c := range cases {
		got := height(c.root)
		if got != c.want {
			t.Errorf("height(%v) == %d, want %d", c.root, got, c.want)
		}
	}
}
