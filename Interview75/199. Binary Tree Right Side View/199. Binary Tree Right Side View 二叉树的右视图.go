package main

//*
// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	record := [][]*TreeNode{}
	record = append(record, []*TreeNode{root})

	for i := 0; i < len(record); i++ {
		current := []*TreeNode{}
		for j := 0; j < len(record[i]); j++ {
			if record[i][j].Left != nil {
				current = append(current, record[i][j].Left)
			}
			if record[i][j].Right != nil {
				current = append(current, record[i][j].Right)
			}
		}
		if len(current) > 0 {
			record = append(record, current)
		}
		res = append(res, record[i][len(record[i])-1].Val)
	}

	return res
}
