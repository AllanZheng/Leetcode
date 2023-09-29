package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return findDepth(root, 0, root.Val)
}

func findDepth(root *TreeNode, result int, maximum int) int {
	if maximum <= root.Val {
		result = result + 1
		maximum = root.Val
	}
	if root.Left != nil {
		result = findDepth(root.Left, result, maximum)
	}
	if root.Right != nil {
		result = findDepth(root.Right, result, maximum)
	}
	return result

}
