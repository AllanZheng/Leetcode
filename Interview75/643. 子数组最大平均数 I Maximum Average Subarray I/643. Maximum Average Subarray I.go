package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	length := len(nums)
	average := 0.0
	maximum := 0.0
	pre := 0
	suf := k - 1
	sum := 0.0
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	average = sum / float64(k)
	maximum = average
	for suf < length-1 {
		suf = suf + 1
		sum = sum - float64(nums[pre]-nums[suf])
		average = sum / float64(k)
		if average > maximum {
			maximum = average
		}
		pre = pre + 1
	}
	return maximum
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3, 9, 1}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}
