package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	record := map[int]int{}
	times := map[int]int{}
	for i := 0; i < len(arr); i++ {
		record[arr[i]]++
	}

	for _, v := range record {
		if times[v] != 0 {
			return false
		} else {
			times[v] = 1
		}
	}
	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
