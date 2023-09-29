package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	mid := make(map[int][][]int, 0)
	res := make([][]int, 0)
	res, _ = backtrack(candidates, target, mid)
	return res
}
func backtrack(candidates []int, target int, mid map[int][][]int) ([][]int, map[int][][]int) {
	res := make([][]int, 0)
	if len(candidates) == 1 {
		if target >= candidates[0] {
			div := target / candidates[0]
			sub := make([]int, 0)
			for i := 1; i <= div; i++ {
				sub = append(sub, candidates[0])
				ele4 := make([]int, len(sub))
				copy(ele4, sub)
				if i*candidates[0] == target {
					res = append(res, ele4)
				} else {
					mid[i*candidates[0]] = append(mid[i*candidates[0]], ele4)
				}
			}
		} else {
			return res, mid
		}
	} else {
		res, mid = backtrack(candidates[1:], target, mid)
		if target >= candidates[0] {
			div := target / candidates[0]
			for j, arr := range mid {
				if target-j >= 0 {
					if (target-j)%candidates[0] == 0 {
						cycle := (target - j) / candidates[0]
						addition := make([]int, 0)
						for k := 1; k <= cycle; k++ {
							addition = append(addition, candidates[0])
							for a := range arr {
								x := append(arr[a], addition...)
								ele := make([]int, len(x))
								copy(ele, x)
								if k < cycle {
									z := candidates[0]*k + j
									mid[z] = append(mid[z], ele)
								} else {

									res = append(res, ele)
								}
							}
						}

					} else {
						sub2 := make([]int, 0)
						for b := candidates[0]; b < target-j; b = b + candidates[0] {
							sub2 = append(sub2, candidates[0])
							for a := range arr {
								y := append(arr[a], sub2...)
								z := b + j
								ele2 := make([]int, len(y))
								copy(ele2, y)
								mid[z] = append(mid[z], ele2)
							}
						}
					}
				}

			}
			sub3 := make([]int, 0)
			for i := 1; i <= div; i++ {
				sub3 = append(sub3, candidates[0])
				ele3 := make([]int, len(sub3))
				copy(ele3, sub3)
				if i*candidates[0] == target {

					res = append(res, ele3)
				} else {

					mid[i*candidates[0]] = append(mid[i*candidates[0]], ele3)
				}
			}
		}

	}
	return res, mid
}
func main() {
	var result [][]int
	//result = combinationSum([]int{7, 3, 2}, 18)
	result = combinationSum([]int{3, 2, 8}, 18)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Print(result[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
