package main

import "fmt"

func distinctDifferenceArray(nums []int) []int {
	res := []int{}
	var prearr, sufarr map[int]int
	for i := 0; i < len(nums); i++ {

		prearr = make(map[int]int, 0)
		sufarr = make(map[int]int, 0)
		for pre := i; pre >= 0; pre-- {
			prearr[nums[pre]] = 1
		}

		for suf := i + 1; suf < len(nums); suf++ {
			sufarr[nums[suf]] = 1
		}

		res = append(res, len(prearr)-len(sufarr))
	}

	return res
}

func main() {
	res := (distinctDifferenceArray([]int{37, 37, 37, 37, 33}))
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
