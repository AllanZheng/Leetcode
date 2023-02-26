# 1305. All Elements in Two Binary Search Trees
# Definition for a binary tree node.
class TreeNode:
     def __init__(self, x):
         self.val = x
         self.left = None
         self.right = None

class Solution:
    def getAllElements(self, root1: TreeNode, root2: TreeNode) -> List[int]:
        arr1 = []
        arr2 = []
        arr1 = Solution.traversal(Solution, root1, arr1)

        arr2 = Solution.traversal(Solution, root2, arr2)
        res = arr1 + arr2
        res.sort()
        return res

    def traversal(self, root: TreeNode, arr: List[int]) -> List[int]:
        if root != None:
            arr.append(root.val)
            if root.right != None:
                arr = Solution.traversal(Solution, root.right, arr)
            if root.left != None:
                arr = Solution.traversal(Solution, root.left, arr)
        return arr
