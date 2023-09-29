package main

import (
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	rec := map[string]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		k := len(nums) - 1
		j := i + 1
		for k > j {
			current := nums[j] + nums[k]
			if current == -nums[i] {
				if _, ok := rec[strconv.Itoa(nums[i])+strconv.Itoa(nums[j])+strconv.Itoa(nums[k])]; !ok {
					res = append(res, []int{nums[i], nums[j], nums[k]})
					rec[strconv.Itoa(nums[i])+strconv.Itoa(nums[j])+strconv.Itoa(nums[k])] = 1
				}
				k--
				j++
			} else if current < -nums[i] {
				k--

			} else {
				j++
			}
		}
	}
	return res
}

func main() {
	input := []int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}
	res := threeSum(input)
	for _, cur := range res {
		println(cur[0], cur[1], cur[2])
	}

}
