package main

import "fmt"

func pivotIndex(nums []int) int {
	presum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		presum = append(presum, presum[i-1]+nums[i])
	}
	length := len(presum)
	if presum[length-1]-presum[0] == 0 {
		return 0
	}
	for i := 1; i < length; i++ {
		if presum[length-1]-presum[i-1]-nums[i] == presum[i-1] {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{0}))
	fmt.Println(pivotIndex([]int{1}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
}
