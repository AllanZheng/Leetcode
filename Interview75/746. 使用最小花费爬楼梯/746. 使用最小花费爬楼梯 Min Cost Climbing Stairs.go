package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	record := []int{0, 0}
	length := len(cost)
	for i := 2; i <= length; i++ {
		if record[i-1]+cost[i-1] < record[i-2]+cost[i-2] {
			record = append(record, record[i-1]+cost[i-1])
		} else {
			record = append(record, record[i-2]+cost[i-2])
		}
	}
	return record[length]
}
func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	fmt.Println(minCostClimbingStairs([]int{1, 2, 3, 4, 5, 15, 7}))
}
