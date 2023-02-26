package main

func removeElement(nums []int, val int) int {
	res := len(nums)

	for i := 0; i <= len(nums); i++ {

		if nums[i] == val {

			nums = append(nums[0:i], nums[i:res]...)
			res = res - 1
			i = i - 1
		}
	}
	return res
}
