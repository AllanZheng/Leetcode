package main

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	max := root.Val
	result := 1
	rec := 1
	currentqueue := []*TreeNode{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		currentqueue = queue
		queue = []*TreeNode{}
		sum := 0
		for i := 0; i < len(currentqueue); i++ {
			sum += currentqueue[i].Val
			if currentqueue[i].Left != nil {
				queue = append(queue, currentqueue[i].Left)
			}
			if currentqueue[i].Right != nil {
				queue = append(queue, currentqueue[i].Right)
			}
		}
		if sum > max {
			max = sum
			result = rec
		}
		rec++
	}
	return result
}
