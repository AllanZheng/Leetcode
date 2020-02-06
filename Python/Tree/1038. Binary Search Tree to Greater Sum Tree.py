# 1038. Binary Search Tree to Greater Sum Tree
# Definition for a binary tree node.
class TreeNode:
     def __init__(self, x):
         self.val = x
         self.left = None
         self.right = None

class Solution:
    def bstToGst(self, root: TreeNode) -> TreeNode:
        cur = 0
        res, cur = Solution.recursive(Solution, root, 0)
        return res

    def recursive(self, root: TreeNode, cur: int):

        if root.right != None:
            root.right, cur = Solution.recursive(Solution, root.right, cur)
        root.val = root.val + cur
        cur = root.val
        if root.left != None:
            root.left, cur = Solution.recursive(Solution, root.left, cur)

        return root, cur

