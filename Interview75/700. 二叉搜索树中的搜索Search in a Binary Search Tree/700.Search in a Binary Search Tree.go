package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	var result *TreeNode
	if root.Val == val {
		result = root
		return result
	}
	if root.Right != nil && result == nil {
		result = searchBST(root.Right, val)
	}
	if root.Left != nil && result == nil {
		result = searchBST(root.Left, val)
	}
	return result
}
