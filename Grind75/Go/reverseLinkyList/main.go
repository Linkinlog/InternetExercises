package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// init current, next, prev nodes
	current := head
	var next, prev *ListNode = nil, nil
	// while current isnt nil, set next -> prev, prev -> current, curr -> next 
	for current != nil {
		next = current.Next

		current.Next = prev
		prev = current
		current = next
	}
	// set head to prev
	head = prev
	return head
}
