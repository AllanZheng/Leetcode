# 1008. Construct Binary Search Tree from Preorder Traversal
# Definition for a binary tree node.
from typing import List
class TreeNode:
    def __init__(self, x):
         self.val = x
         self.left = None
         self.right = None

class Solution:
    def bstFromPreorder(self, preorder: List[int]) -> TreeNode:
        res = TreeNode(preorder[0])
        for i in range(1, len(preorder)):
            res = Solution.recurrsive(Solution, preorder[i], res)
        return res

    def recurrsive(self, inpu: int, tree: TreeNode) -> TreeNode:
        if tree == None:
            tree = TreeNode(inpu)
        elif tree.val > inpu:
            tree.left = Solution.recurrsive(Solution, inpu, tree.left)
        else:
            tree.right = Solution.recurrsive(Solution, inpu, tree.right)

        return tree
