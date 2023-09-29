package main

import (
	"fmt"
	"sort"
	"strconv"
)

func combinationSum3(k int, n int) [][]int {
	cur := []int{}
	res := [][]int{}
	res = cal(k, n, cur, res)
	record := map[string]bool{}
	final := [][]int{}
	for i := 0; i < len(res); i++ {
		curstr := ""
		for j := 0; j < k; j++ {
			curstr = strconv.Itoa(res[i][j]) + curstr
		}

		if !record[curstr] {
			record[curstr] = true
			final = append(final, res[i])
		}
	}
	return final

}
func cal(k int, n int, cur []int, res [][]int) [][]int {
	if n <= 0 {
		return res
	}
	if k == 1 {
		if n > 9 {
			return res
		}
		exist := false
		for j := 0; j < len(cur); j++ {
			if cur[j] == n {
				exist = true
				break
			}
		}
		if !exist {
			cur := append(cur, n)
			newcur2 := make([]int, len(cur))
			copy(newcur2, cur)
			sort.Ints(newcur2)
			res = append(res, newcur2)
		}
		return res
	}
	maxi := 0
	if n > 9 {
		maxi = 9
	} else {
		maxi = n - 1
	}
	for i := 1; i <= maxi; i++ {
		exist := false
		for j := 0; j < len(cur); j++ {
			if cur[j] == i {
				exist = true
				break
			}
		}
		if !exist {
			newcur := append(cur, i)
			res = cal(k-1, n-i, newcur, res)

		}
	}
	return res
}
func main() {
	result := combinationSum3(8, 36)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Print(result[i][j])
		}
		fmt.Println()
	}
}
