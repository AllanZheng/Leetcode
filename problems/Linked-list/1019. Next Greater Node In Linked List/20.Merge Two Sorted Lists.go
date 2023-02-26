package Medium

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			return res.Next
		}
		if list2 == nil {
			cur.Next = list1
			return res.Next
		}
		if list1.Val < list2.Val {
			cur.Next = &ListNode{list1.Val, nil}

			list1 = list1.Next

		} else {
			cur.Next = &ListNode{list2.Val, nil}

			list2 = list2.Next

		}
		cur = cur.Next

	}
	return res.Next
}
