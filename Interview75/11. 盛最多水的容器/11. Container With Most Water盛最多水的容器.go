package main

import "fmt"

func maxArea(height []int) int {
	maximum := 0
	i := 0
	j := len(height) - 1
	for i < j {
		if height[i] < height[j] {
			if maximum < (j-i)*height[i] {
				maximum = (j - i) * height[i]
			}
			i++
		} else {
			if maximum < (j-i)*height[j] {
				maximum = (j - i) * height[j]
			}
			j--
		}
	}
	return maximum
}

func maxArea1(height []int) int {
	maximum := 0
	for i := 0; i < len(height); i++ {
		if height[i] != 0 {
			for j := len(height) - 1; j > i; j-- {
				if height[j] >= height[i] {
					if maximum < height[i]*(j-i) {
						maximum = height[i] * (j - i)

					}
					break
				}
				if maximum < height[j]*(j-i) {
					maximum = height[j] * (j - i)

				}
			}
		}
	}
	return maximum
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{0, 1}))
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 25, 7}))
}
