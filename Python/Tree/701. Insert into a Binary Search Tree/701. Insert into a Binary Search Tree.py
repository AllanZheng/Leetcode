# Definition for a binary tree node.
class TreeNode:
     def __init__(self, x):
         self.val = x
         self.left = None
         self.right = None

class Solution:
        def insertIntoBST(self, root: TreeNode, val: int) -> TreeNode:
            stack = []
            stack.append(root)
            while len(stack) != 0:
                cur = stack.pop()
                if val < cur.val:
                    if cur.left != None:
                        stack.append(cur.left)
                    else:
                        newnode = TreeNode(val)
                        cur.left = newnode
                        break
                else:
                    if cur.right != None:
                        stack.append(cur.right)
                    else:
                        newnode = TreeNode(val)
                        cur.right = newnode
                        break
            return root