package main

import "testing"

func TestMiddleNode(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		input *ListNode
		want  *ListNode
	}{
		"case 1": {input: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, want: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}},
		"case 2": {input: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}, want: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := middleNode(tc.input)
			if got == nil {
				t.Fatalf("got: %v, want: %v", got, tc.want)
			}
			if got.Val != tc.want.Val {
				t.Errorf("got: %v, want: %v", got.Val, tc.want.Val)
			}
		})
	}
}
