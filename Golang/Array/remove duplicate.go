package main

func removeDuplicates(nums []int) int {
	res := len(nums)

	for i := 1; i < res; i++ {

		if nums[i] == nums[i-1] {

			nums = append(nums[0:i], nums[i+1:res]...)
			res = res - 1
			i = i - 1
		}
	}
	return res
}
