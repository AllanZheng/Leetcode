package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	var second, start *ListNode
	if head.Next == nil {
		return head
	}
	start = head
	second = head.Next

	for second.Next != nil {
		second = second.Next
		if second.Next != nil {
			second = second.Next
		} else {
			break
		}
		start = start.Next
	}
	be := start.Next
	start.Next = be.Next
	return head
}
