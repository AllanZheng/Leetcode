package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return findDepth(root, 0)
}

func findDepth(root *TreeNode, depth int) int {
	depth++
	left := depth
	right := depth
	if root.Left != nil {
		left = findDepth(root.Left, depth)
	}
	if root.Right != nil {
		right = findDepth(root.Right, depth)
	}
	if left > right {
		return left
	} else {
		return right
	}

}
