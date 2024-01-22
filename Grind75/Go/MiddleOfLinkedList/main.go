package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	current := head
	sum := 0
	for current != nil {
		sum += 1
		current = current.Next
	}

	half := sum / 2

	halfNode := head
	for i := 0; i < half; i++ {
		halfNode = halfNode.Next
	}

	return halfNode
}
