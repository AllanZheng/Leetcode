package main

import "fmt"

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	record := []int{0, 1, 1}
	for i := 2; i <= n; i++ {
		record = append(record, record[i]+record[i-1]+record[i-2])
	}
	return record[n]
}
func main() {
	fmt.Println(tribonacci(25))
}
