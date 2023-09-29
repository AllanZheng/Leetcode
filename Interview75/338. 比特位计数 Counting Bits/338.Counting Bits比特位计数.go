package main

import "fmt"

func countBits(n int) []int {
	result := []int{0}
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			result = append(result, result[i/2])
		} else {
			result = append(result, result[i-1]+1)
		}
	}
	return result
}

func countBits2(n int) []int {
	result := []int{0}
	for i := 1; i <= n; i++ {
		result = append(result, result[i&(i-1)]+1)
	}
	return result
}

func main() {
	fmt.Println(countBits2(10))
}
