package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	maximum := 0
	var newList *ListNode

	cur := head
	for cur != nil {
		newNode := ListNode{
			Val:  cur.Val,
			Next: newList,
		}
		newList = &newNode
		cur = cur.Next
	}

	// newNode := ListNode{
	// 	Val:  cur.Val,
	// 	Next: newList,
	// }
	// newList.Next = &newNode
	cur = head
	for cur != nil {
		value := 0
		if newList != nil {
			value = cur.Val + newList.Val
			newList = newList.Next
		}
		if value > maximum {
			maximum = value
		}

		cur = cur.Next

	}
	return maximum
}

func main() {
	LinkedList := ListNode{
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
	}
	fmt.Println(pairSum(&LinkedList))
}
