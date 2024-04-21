package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//深搜解法，还有动态规划解法
func longestZigZag(root *TreeNode) int {
	maxleft, maxright := 0, 0
	if root.Right != nil {
		maxright = searchright(root.Right, maxright)
	}
	if root.Left != nil {
		maxleft = searchleft(root.Left, maxleft)
	}
	return max(maxright, maxleft)
}

func searchleft(node *TreeNode, maximum int) int {
	maximum = maximum + 1
	if node.Right != nil {
		maximum = searchright(node.Right, maximum)
	}
	init := 0
	if node.Left != nil {
		init = searchleft(node.Left, init)
	}
	return max(maximum, init)
}

func searchright(node *TreeNode, maximum int) int {
	maximum = maximum + 1
	if node.Left != nil {
		maximum = searchleft(node.Left, maximum)
	}
	init := 0
	if node.Right != nil {
		init = searchright(node.Right, init)
	}
	return max(maximum, init)
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
