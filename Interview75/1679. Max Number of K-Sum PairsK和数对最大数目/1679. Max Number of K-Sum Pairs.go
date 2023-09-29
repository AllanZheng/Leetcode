package main

import (
	"fmt"
	"sort"
)

func maxOperations2(nums []int, k int) int {
	result := 0
	record := map[int]int{}
	for i := 0; i < len(nums); i++ {
		record[nums[i]]++
	}
	for key, v := range record {
		if record[k-key] > 0 {
			if v < record[k-key] {
				result = result + v
			} else {
				result = result + record[k-key]
			}

		}
	}

	return result / 2
}

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	ans := 0
	for left < right {
		temp := nums[left] + nums[right]
		if temp < k {
			left++
		} else if temp > k {
			right--
		} else {
			ans++
			left++
			right--
		}
	}
	return ans
}

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))
	fmt.Println(maxOperations([]int{1, 1, 1, 1, 1, 1}, 2))
	fmt.Println(maxOperations([]int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}, 2))
}
