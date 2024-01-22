package main

import "testing"

func TestReverseList(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input *ListNode
		want  *ListNode
	}{
		"Case 1": {
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
		"Case 2": {
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
		"Case 3": {
			input: &ListNode{},
			want:  &ListNode{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := reverseList(tc.input)
			if got.Val != tc.want.Val {
				t.Errorf("got %d want %d", got.Val, tc.want.Val)
			}
		})
	}
}
