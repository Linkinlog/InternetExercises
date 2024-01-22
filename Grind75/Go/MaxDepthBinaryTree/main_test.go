package main

import "testing"

func TestMaxDepth(t *testing.T) {
	cases := map[string]struct {
		input *TreeNode
		want  int
	}{
		"Case 1": {input: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, want: 3},
		"Case 2": {input: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, want: 2},
		"Case 3": {input: &TreeNode{Val: 0}, want: 1},
		"Case 4": {input: nil, want: 0},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := maxDepth(tc.input)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
