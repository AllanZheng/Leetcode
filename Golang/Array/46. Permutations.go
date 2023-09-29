package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	for i := 0; i < len(nums); i++ {
		curres := backtrack(nums, [][]int{[]int{nums[i]}})
		for j := 0; j < len(curres); j++ {
			res = append(res, curres[i])
		}
	}

	return res
}

func backtrack(nums []int, res [][]int) [][]int {
	if len(nums) == 1 {
		for i := 0; i < len(res); i++ {
			res[i] = append(res[i], nums[i])
		}
		return res
	} else {
		for i := 0; i < len(nums); i++ {
			var cur [][]int
			for j := 0; j < len(res); j++ {
				cur = append(cur, append(res[j], nums[i]))
			}
			return backtrack(append(nums[0:i], nums[i+1:]...), cur)
		}

	}
	return res
}

/*
func permute1(nums []int) [][]int {

	res := make([][]int, math.po)
	return res
}
*/

func main() {
	result := permute([]int{1, 2, 3})
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
