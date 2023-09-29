package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	record1 := map[int]int{}
	record2 := map[int]int{}
	result := [][]int{[]int{}, []int{}}
	for i := 0; i < len(nums1); i++ {
		record1[nums1[i]] = 1
	}
	for i := 0; i < len(nums2); i++ {
		record2[nums2[i]] = 1
	}
	for i := 0; i < len(nums1); i++ {
		if record2[nums1[i]] != 1 {
			result[0] = append(result[0], nums1[i])
			record2[nums1[i]] = 1
		}
	}
	for i := 0; i < len(nums2); i++ {
		if record1[nums2[i]] != 1 {
			result[1] = append(result[1], nums2[i])
			record1[nums2[i]] = 1
		}
	}
	return result
}
func printResult(input [][]int) {
	for i := 0; i < 2; i++ {
		for j := 0; j < len(input[i]); j++ {
			fmt.Print(input[i][j], " ")
		}
		fmt.Println()
	}
}
func main() {
	result := findDifference([]int{1, 2, 3}, []int{2, 4, 6})
	printResult(result)
	result1 := findDifference([]int{1, 2, 3}, []int{4, 5, 6, 4, 5, 6})
	printResult(result1)
}
