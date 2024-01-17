package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// 	Traverse the list individually and keep putting the node addresses in a Hash Table.
	// At any point, if NULL is reached then return false
	// If the next of the current nodes points to any of the previously stored nodes in  Hash then return true.
	if head == nil {
		return false
	}
	h := map[*ListNode]bool{head: true}
	cur := head
	for cur != nil && cur.Next != nil {
		if _,ok := h[cur]; ok {
			return true
		} else {
			h[cur] = true
		}
		cur = cur.Next
	}
	return false
}
