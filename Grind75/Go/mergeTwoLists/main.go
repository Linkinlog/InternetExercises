package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	returnList := &ListNode{}
	current := returnList
loopy:
	for {
		switch true {
		case list1 == nil, list2 == nil, list1.Next == nil && list2.Next == nil:
			break loopy
		case list1.Val > list2.Val:
			// list2.val goes onto returnList
			current.Next = list2
			current = current.Next
			list2 = list2.Next
		case list1.Val < list2.Val, list1.Val == list2.Val:
			// list1.val goes onto returnList
			current.Next = list1
			current = current.Next
			list1 = list1.Next
		}
	}

	return returnList.Next
}
