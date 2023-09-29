package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	length := len(temperatures)
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		temp := temperatures[i]
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}
func main() {
	result := dailyTemperatures([]int{1, 2, 3, 4, 5})
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
