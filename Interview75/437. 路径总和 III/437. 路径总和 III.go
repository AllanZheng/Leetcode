package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSumSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	result := 0
	if root.Val == targetSum {
		result++
	}
	result = result + pathSumSum(root.Left, targetSum-root.Val)
	result = result + pathSumSum(root.Right, targetSum-root.Val)
	return result
}

func pathSum(root *TreeNode, targetSum int) int {

	if root == nil {
		return 0
	}
	return pathSumSum(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}
