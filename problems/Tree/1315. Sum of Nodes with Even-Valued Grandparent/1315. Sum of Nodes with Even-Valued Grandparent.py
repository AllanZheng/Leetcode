# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def sumEvenGrandparent(self, root: TreeNode) -> int:
        res = 0
        ressum = []
        ressum = Solution.recurrison(Solution, root, ressum)
        for i in ressum:
            res += i
        return res

    def recurrison(self, root: TreeNode, res: List[int]) -> List[int]:
        cur = 0

        if root.val % 2 == 0:
            if root.right != None:
                if root.right.right != None:
                    cur += root.right.right.val

                if root.right.left != None:
                    cur += root.right.left.val

            if root.left != None:
                if root.left.right != None:
                    cur += root.left.right.val

                if root.left.left != None:
                    cur += root.left.left.val

        res.append(cur)

        if root.right != None:
            res = Solution.recurrison(Solution, root.right, res)
        if root.left != None:
            res = Solution.recurrison(Solution, root.left, res)
        return res


