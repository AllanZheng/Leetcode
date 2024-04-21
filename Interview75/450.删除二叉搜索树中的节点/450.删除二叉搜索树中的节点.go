package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	cur := root
	for root != nil && root.Val != key {
		if root.Val < key {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	if root == nil {
		return nil
	}
	for root.Right != nil {
		next := root.Right
		if next.Left != nil {
			find(cur.Right, next.Left)
		}
		root = next
		return cur

	}
	root = root.Left
	return cur
}
func find(root *TreeNode, left *TreeNode) {
	if root.Left != nil {
		find(root.Right, left)
	} else {
		root.Left = left
	}
}
