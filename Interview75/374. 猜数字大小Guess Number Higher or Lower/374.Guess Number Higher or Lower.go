package main

import "fmt"

func guess(num int) int {
	pick := 8
	if num == pick {
		return 0
	} else if num < pick {
		return 1
	} else {
		return -1
	}

}

func guessNumber(n int) int {

	low := 1
	high := n
	for low <= high {

		current := low + (high-low)/2
		res := guess(current)
		if res == 0 {
			return current
		} else if res > 0 {
			low = current + 1

		} else {
			high = current
		}
	}
	return -1
}

func main() {
	fmt.Println(guessNumber(10))
}
