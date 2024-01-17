package main

import "testing"

func newListNode(val int) *ListNode {
	return &ListNode{val, nil}
}

func (l *ListNode) append(ln *ListNode) {
	l.Next = ln
}


func TestHasCycle(t *testing.T) {
	tests := map[string]struct {
		input func() *ListNode
		want  bool
	}{
		"1": {
			input: func() *ListNode {
				l1 := newListNode(3)
				l2 := newListNode(2)
				l3 := newListNode(0)
				l4 := newListNode(-4)
				l1.append(l2)
				l2.append(l3)
				l3.append(l4)
				l4.append(l2)
				return l1
			},
			want: true,
		},
		"2": {
			input: func() *ListNode {
				l1 := newListNode(1)
				l2 := newListNode(2)
				l1.append(l2)
				l2.append(l1)
				return l1
			},
			want: true,
		},
		"3": {
			input: func() *ListNode {
				l1 := newListNode(1)
				return l1
			},
			want: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := hasCycle(tc.input()); got != tc.want {
				t.Fatalf("hasCycle(%v) = %v, want %v", tc.input(), got, tc.want)
			}
		})
	}
} 

