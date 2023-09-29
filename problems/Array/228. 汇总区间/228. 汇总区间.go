package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := []string{}

	length := len(nums)
	if length == 0 {
		return result
	}
	if length == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	start := nums[0]
	end := nums[0]
	dir := "->"
	for i := 1; i <= length; i++ {
		if i == length || nums[i]-1 != nums[i-1] {
			if start == end {
				result = append(result, strconv.Itoa(nums[i-1]))

			} else {
				current := strconv.Itoa(start) + dir + strconv.Itoa(end)
				result = append(result, current)

			}
			if i < length {
				start = nums[i]
				end = nums[i]
			}

		} else {
			end = nums[i]
		}
	}

	return result
}
func main() {
	res := summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})
	//res1 := summaryRanges([]int{0, 1, 2, 4, 5, 7})
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

}
