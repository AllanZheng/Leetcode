# Definition for singly-linked list.
class ListNode:
     def __init__(self, x):
         self.val = x
         self.next = None

class Solution:
    def getDecimalValue(self, head: ListNode) -> int:
        stack = []
        res = 0
        weight = 1
        while head.next != None:
            stack.append(head.val)
            head = head.next
        while len(stack) !=0:
            cur = stack.pop()
            res = res + weight*cur
            weight = weight * 2
        return res
# print(Solution.getDecimalValue(Solution, ListNode{ListNode.val: 1, next: ListNode{ListNode.val: 0, next: ListNode{ListNode.val: 1, next: None}}}))