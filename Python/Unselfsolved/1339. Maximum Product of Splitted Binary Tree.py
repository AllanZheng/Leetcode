#1339. Maximum Product of Splitted Binary Tree
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
         self.val = x
         self.left = None
         self.right = None

class Solution:
    def maxProduct(self, root: TreeNode) -> int:
        sumarr = []
        def recursive(Node):
            if Node == None:
                return 0
            allsum = Node.val + recursive(Node.right)+ recursive(Node.left)
            sumarr.append(allsum)
            return allsum
        total =  recursive(root)
        maximum = 0

        for i in sumarr:
            cur = i *(total-i)
            if cur>maximum:
                maximum = cur
        return maximum%(10 ** 9 + 7)