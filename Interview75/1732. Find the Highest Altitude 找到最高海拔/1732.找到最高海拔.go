package main

import "fmt"

func largestAltitude(gain []int) int {
	max, cur := 0, 0
	for i := 0; i < len(gain); i++ {
		cur = cur + gain[i]
		if max < cur {
			max = cur
		}
	}
	return max
}

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
