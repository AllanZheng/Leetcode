package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	var odd, even *ListNode
	odd = head
	if odd == nil {
		return nil
	}
	even = head.Next
	if even == nil {
		return odd
	}
	evenhead := even
	for even != nil {
		odd.Next = even.Next
		if odd.Next == nil {
			odd.Next = evenhead
			break
		}
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	if odd.Next != evenhead {
		odd.Next = evenhead
	}
	return head
}
func main() {
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := oddEvenList(head1)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
