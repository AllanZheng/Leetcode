package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pparents, qparents := []*TreeNode{}, []*TreeNode{}
	pparents = find(root, p)
	qparents = find(root, q)
	lenp := len(pparents)
	lenq := len(qparents)
	length := 0
	sub := 0
	if lenp < lenq {
		length = lenq
		sub = lenq - lenp
		for i := 0; i < length; i++ {
			if pparents[i] == qparents[i+sub] {
				return pparents[i]
			}
		}
	} else {
		length = lenp
		sub = lenp - lenq
		for i := 0; i < length; i++ {
			if pparents[i+sub] == qparents[i] {
				return qparents[i]
			}
		}
	}

	return nil
}

func find(root, target *TreeNode) []*TreeNode {

	if root == target {
		return []*TreeNode{target}
	}
	if root.Left != nil {
		cur := find(root.Left, target)
		if cur != nil {
			cur = append(cur, root)
			return cur
		}
	}
	if root.Right != nil {
		cur := find(root.Right, target)
		if cur != nil {
			cur = append(cur, root)
			return cur
		}
	}
	return nil
}
