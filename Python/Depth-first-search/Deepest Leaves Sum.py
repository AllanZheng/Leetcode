# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
         self.val = x
         self.left = None
         self.right = None

class Solution:
    def deepestLeavesSum(self, root: TreeNode) -> int:
        stack = []
        res = 0
        curdepth = 0
        maxdepth = 0
        stack.append(root)
        while len(stack) > 0:
            if root.left != None:
                if root.left.val == -1:
                    root.left = None

                else:
                    stack.append(root)
                    curdepth = curdepth + 1
                    root = root.left
                continue
            elif root.right != None:

                if root.right.val == -1:
                    root.right = None

                else:
                    stack.append(root)
                    curdepth = curdepth + 1
                    root = root.right
                continue
            if curdepth > maxdepth:
                res = root.val
                maxdepth = curdepth
            elif curdepth == maxdepth:
                res = res + root.val
            root.val = -1
            root = stack.pop()
            curdepth = curdepth - 1

        return res

# print(Solution.deepestLeavesSum(Solution,TreeNode))

