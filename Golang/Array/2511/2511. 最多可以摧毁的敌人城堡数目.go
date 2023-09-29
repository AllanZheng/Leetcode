package main

import "fmt"

func captureForts(forts []int) int {
	maximum := 0
	for i := 0; i < len(forts); i++ {
		if forts[i] == 1 {
			front := i - 1
			end := i + 1
			for front >= 0 {
				if forts[front] == 1 {
					break
				} else if forts[front] == -1 {
					if i-front-1 > maximum {
						maximum = i - front - 1

					}
					break
				} else {
					front--
				}
			}
			for end < len(forts) {
				if forts[end] == 1 {
					break
				} else if forts[end] == -1 {
					if end-i-1 > maximum {
						maximum = end - i - 1

					}
					break
				} else {
					end++
				}
			}
		}
	}
	return maximum
}

func main() {
	// fmt.Println(captureForts([]int{1, 0, 0, -1, 0, 0, 0, 0, 1}))
	// fmt.Println(captureForts([]int{0, 0, 1, -1}))

	fmt.Println(captureForts([]int{-1, 0, 1, 0, 0, 0}))
}
