package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maximum := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > maximum {
			maximum = candies[i]
		}
	}
	result := []bool{}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maximum {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}
func main() {
	res := kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
