package Medium
/**
 * Definition for singly-linked list.
 */
  type ListNode struct {
      Val int
      Next *ListNode
 }

func nextLargerNodes(head *ListNode) []int {

	res := []int{}
	for head != nil {
		curmax := head
		for curmax!=nil&&curmax.Val<=head.Val{
   		curmax = curmax.Next
		}
		if curmax==nil{
			res =append(res,0)
		}else{
			res= append(res,curmax.Val)
		}
		head = head.Next
	}
	return res

}