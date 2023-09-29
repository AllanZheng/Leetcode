package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	pre := head
	if head == nil {
		return nil
	}
	head = head.Next
	pre.Next = nil

	for head != nil {
		cur := head
		head = head.Next
		cur.Next = pre
		pre = cur
	}

	return pre
}

func main() {
	//head := &ListNode{1, &ListNode{3, &ListNode{2, &ListNode{4, nil}}}}
	head1 := &ListNode{1, nil}
	result := reverseList(head1)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
