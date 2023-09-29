package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := []int{}
	var sum int64
	sum = int64(nums[0])

	zero := 0
	if sum == 0 {
		zero = 1
		sum = 1
	}
	length := len(nums)
	for i := 1; i < length; i++ {
		if nums[i] != 0 || zero == 1 {
			sum = sum * int64(nums[i])
		} else {
			zero = 1
		}

	}
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			if zero == 0 {
				result = append(result, int(sum/int64(nums[i])))
			} else {
				result = append(result, 0)
			}
		} else {
			result = append(result, int(sum))
		}
	}
	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
	fmt.Println(productExceptSelf([]int{2, 3, 0, 0}))
	fmt.Println(productExceptSelf([]int{0, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{0, 0, 3, 4}))
}
