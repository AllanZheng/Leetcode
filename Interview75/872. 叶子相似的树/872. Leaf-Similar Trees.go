package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	root1leafs := []int{}
	root2leafs := []int{}
	var dfs func(*TreeNode, []int) []int
	dfs = func(node *TreeNode, res []int) []int {
		if node.Left != nil {
			res = dfs(node.Left, res)
		}
		if node.Right != nil {
			res = dfs(node.Right, res)
		}
		if node.Left == nil && node.Right == nil {
			res = append(res, node.Val)
		}
		return res
	}
	root1leafs = dfs(root1, root1leafs)
	root2leafs = dfs(root2, root2leafs)
	if len(root1leafs) != len(root2leafs) {
		return false
	}

	for i := 0; i < len(root1leafs); i++ {
		if root1leafs[i] != root2leafs[i] {
			return false
		}
	}
	return true
}

func main() {
	root1 := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	root2 := TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}
	root3 := TreeNode{1, nil, nil}
	root4 := TreeNode{2, nil, nil}
	fmt.Println(leafSimilar(&root1, &root2))
	fmt.Println(leafSimilar(&root3, &root4))
}
