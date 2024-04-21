// ; 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

// ; candidates 中的每个数字在每个组合中只能使用 一次 。
// ; 注意：解集不能包含重复的组合。
// ; 示例 1:
// ; 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// ; 输出:
// ; [[1,1,6],[1,2,5],[1,7],[2,6]]
// ; 示例 2:

// ; 输入: candidates = [2,5,2,1,2], target = 5,
// ; 输出:
// ; [[1,2,2],[5]]

// ; 提示:
// ; 1 <= candidates.length <= 100
// ; 1 <= candidates[i] <= 50
// ; 1 <= target <= 30
package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var frequency [][2]int
	for _, num := range candidates {
		if frequency == nil || num != frequency[len(frequency)-1][0] {
			frequency = append(frequency, [2]int{num, 1})

		} else {
			frequency[len(frequency)-1][1]++
		}
	}
	var seq []int
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), seq...))
			return
		}
		if pos == len(frequency) || rest < frequency[pos][0] {
			return
		}

		dfs(pos+1, rest)
		most := min(rest/frequency[pos][0], frequency[pos][1])
		for i := 1; i <= most; i++ {
			seq = append(seq, frequency[pos][0])
			dfs(pos+1, rest-i*frequency[pos][0])
		}
		seq = seq[:len(seq)-most]
	}
	dfs(0, target)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	cand := []int{10, 1, 2, 7, 6, 1, 5}
	answer := [][]int{}
	target := 8
	answer = combinationSum2(cand, target)
	for i := 0; i < len(answer); i++ {
		for j := 0; j < len(answer[i]); j++ {
			fmt.Print(answer[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}

}
