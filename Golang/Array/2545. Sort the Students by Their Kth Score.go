package main

import (
	"fmt"
	"sort"
)

func sortTheStudents(score [][]int, k int) [][]int {
	record := make(map[int][]int)
	arr := []int{}
	var result [][]int
	for i := 0; i < len(score); i++ {
		record[score[i][k]] = score[i]
		arr = append(arr, score[i][k])
	}
	sort.Ints(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, record[arr[i]])
	}
	return result
}

func main() {
	martix := [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}
	k := 2
	res := sortTheStudents(martix, k)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Print(res[i][j])
			fmt.Print(" ")

		}
		fmt.Println()
	}
}
