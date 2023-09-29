package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	first := nums[0]
	second := math.MaxInt32
	length := len(nums)
	if length < 3 {
		return false
	}
	for i := 1; i < length; i++ {
		if nums[i] > second {
			return true
		} else if nums[i] > first {
			second = nums[i]

		} else {
			first = nums[i]
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	fmt.Println(increasingTriplet([]int{3, 2, 1}))

}
