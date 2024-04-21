package main

import (
	"fmt"
	"sort"
)

// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

// candidates 中的每个数字在每个组合中只能使用 一次 。
// 注意：解集不能包含重复的组合。
// 示例 1:
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 输出:
// [[1,1,6],[1,2,5],[1,7],[2,6]]
// 示例 2:

// 输入: candidates = [2,5,2,1,2], target = 5,
// 输出:
// [[1,2,2],[5]]

// 提示:
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30

func main() {
	cand := []int{10, 1, 2, 7, 6, 1, 5}
	answer := [][]int{}
	target := 8
	sort.Ints(cand)
	for i := 0; i < len(cand); i++ {
		if cand[i] == target {
			answer = append(answer, []int{cand[i]})
		} else if cand[i] > target {
			continue
		} else {
			result := []int{cand[i]}
			result = find(target-cand[i], result, cand, i+1)
			if len(result) == 0 {
				continue
			} else {
				answer = append(answer, result)
			}
		}
	}
	for i := 0; i < len(answer); i++ {
		for j := 0; j < len(answer[i]); j++ {
			fmt.Print(answer[i][j])
		}
		fmt.Println()
	}
}

func find(target int, result []int, cand []int, current_pos int) []int {
	if cand[current_pos] > target {
		return []int{}
	} else if cand[current_pos] == target {
		result = append(result, cand[current_pos])
		return result
	} else {
		result = append(result, cand[current_pos])
		return find(target-cand[current_pos], result, cand, current_pos+1)
	}
}
